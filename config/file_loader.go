package config

import (
    "os"
    "os/user"
    "encoding/json"
    "io/ioutil"
    "log"
    "path/filepath"
)

const (
    localConfigurationFile = ".hermes.config.json"
)

// Config struct with configuration data for the client application
type Config struct {
    AmqpUrl string `json:"hermes_amqp"`
    ClientHandlerServer string `json:"hermes_client_handler_server"`
}

// LoadConfig loads the globalConfig structure from a JSON-based stream
func LoadConfig() Config {
    var globalConfig Config
    var dataStream []byte
    var err error

    usr, _ := user.Current()
    fullPathLocalConfigurationFile, _ := filepath.Abs(filepath.Join(usr.HomeDir, localConfigurationFile))
    if _, jErr := os.Stat(fullPathLocalConfigurationFile); jErr == nil {
        // "~/.hermes.config.json" exists
        dataStream, err = ioutil.ReadFile(fullPathLocalConfigurationFile)
        if err != nil {
            panic(err)
        }
    } else {
        // no configuration option available
        log.Fatal("> No configuration option is available; fatal")
    }

    err = json.Unmarshal([]byte(dataStream), &globalConfig)
    if err != nil {
        panic(err)
    }

    return globalConfig
}
