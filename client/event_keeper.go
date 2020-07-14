package client

import (
    "log"
)

func SetupClient() {
    err := setupConsumer()
    if err != nil {
        log.Fatal("[CLIENT][AMQP] Failed setup client; fatal")
    }
}
