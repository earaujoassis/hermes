package tunnel

import (
    "bytes"

    "github.com/streadway/amqp"

    "github.com/earaujoassis/hermes/server/models"
    "github.com/earaujoassis/hermes/config"
)

const (
    CorrelationIdSize int = 32
)

func dispatchRequest(requestBuffer []byte) (bytes.Buffer, error) {
    responseBuffer := &bytes.Buffer{}

    cfg := config.CreateTLSConfig(config.GetEnvVar("HERMES_CACERTFILE"),
        config.GetEnvVar("HERMES_CERTFILE"),
        config.GetEnvVar("HERMES_KEYFILE"))
    conn, err := amqp.DialTLS(config.GetEnvVar("HERMES_AMQP"), cfg)
    if err != nil {
        return *responseBuffer, err
    }
    defer conn.Close()

    channel, err := conn.Channel()
    if err != nil {
        return *responseBuffer, err
    }
    defer channel.Close()

    queue, err := channel.QueueDeclare("", false, false, true, false, nil)
    if err != nil {
        return *responseBuffer, err
    }

    msgs, err := channel.Consume(queue.Name, "", true, false, false, false, nil)
    if err != nil {
        return *responseBuffer, err
    }

    correlationId := models.GenerateRandomString(CorrelationIdSize)
    err = channel.Publish("", "requests", false, false, amqp.Publishing {
        ContentType:   "text/plain",
        CorrelationId: correlationId,
        ReplyTo:       queue.Name,
        Body:          requestBuffer,
    })
    if err != nil {
        return *responseBuffer, err
    }

    for d := range msgs {
        if correlationId == d.CorrelationId {
                responseBuffer.Write([]byte(d.Body))
                break
        }
    }

    return *responseBuffer, nil
}
