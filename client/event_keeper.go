package client

import (
    "log"
)

func SetupClient() {
    err := setupConsumer()
    if err != nil {
        log.Fatal("[CLIENT][AMQP] Panic: failed to setup client: ", err.Error())
    }
}
