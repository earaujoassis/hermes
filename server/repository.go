package server

import (
    "fmt"
    "github.com/jinzhu/gorm"
    // Uses Postgres for GORM setup
    _ "github.com/jinzhu/gorm/dialects/postgres"

    "github.com/earaujoassis/hermes/server/models"
    "github.com/earaujoassis/hermes/server/config"
)

var dataStore *gorm.DB

// Start is used to setup the models within the application
func RepositoryStart() {
    GetDataStoreConnection().AutoMigrate(&models.Client{})
}

// GetDataStoreConnection is used to obtain a connection with
//      the Postgres datastore
func GetDataStoreConnection() *gorm.DB {
    if dataStore != nil {
        return dataStore
    }
    var err error
    var databaseName = fmt.Sprintf("%v_%v",
        config.GetEnvVarDefault("HERMES_DATASTORE_PREFIX", "hermes"),
        config.GetEnvVarDefault("HERMES_ENV", "development"))
    var databaseConnectionData = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
        config.GetEnvVarDefault("HERMES_DATASTORE_HOST", "localhost"),
        config.GetEnvVarDefault("HERMES_DATASTORE_USERNAME", "postgres"),
        databaseName,
        config.GetEnvVarDefault("HERMES_DATASTORE_SSL_MODE", "disable"),
        config.GetEnvVarDefault("HERMES_DATASTORE_PASSWORD", ""),
    )
    fmt.Printf("Connected to the following datastore: %v\n", databaseConnectionData)
    dataStore, err = gorm.Open("postgres", databaseConnectionData)
    if err != nil {
        panic(fmt.Sprintf("Failed to connect datastore: %v\n", err))
    }
    return dataStore
}
