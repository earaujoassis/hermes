package config

import (
    "os"
    "encoding/json"
    "io/ioutil"
    "log"
)

var globalConfig Config

// Config struct with configuration data for the client application
type Config struct {
    AmqpUrl string `json:"hermes_amqp"`
    ClientHandlerServer string `json:"hermes_client_handler_server"`
    CACertFile string `json:"hermes_cacertfile"`
    CertFile string `json:"hermes_certfile"`
    KeyFile string `json:"hermes_keyfile"`
}

// GetGlobalConfig returns the global configuration struct for the application
func GetGlobalConfig() Config {
    return globalConfig
}

// LoadConfig loads the globalConfig structure from a JSON-based stream
func LoadConfig(filepath string) Config {
    var dataStream []byte
    var err error

    if _, jErr := os.Stat(filepath); jErr == nil {
        // if "filepath" exists
        dataStream, err = ioutil.ReadFile(filepath)
        if err != nil {
            panic(err)
        }
    } else {
        // otherwise, no configuration option available
        log.Fatal("> No configuration option is available; panic")
    }

    err = json.Unmarshal([]byte(dataStream), &globalConfig)
    if err != nil {
        panic(err)
    }

    return globalConfig
}
