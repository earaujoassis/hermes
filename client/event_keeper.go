package client

import (
    "log"

    "github.com/earaujoassis/hermes/config"
)

func SetupClient(filepath string) {
    config.LoadConfig(filepath)
    err := setupConsumer()
    if err != nil {
        log.Fatal("[CLIENT][AMQP] Panic: failed to setup client: ", err.Error())
    }
}
