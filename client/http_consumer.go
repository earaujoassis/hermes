package client

import (
    "log"

    "github.com/streadway/amqp"

    "github.com/earaujoassis/hermes/config"
)

func setupConsumer() error {
    var globalConfig config.Config = config.GetGlobalConfig()

    cfg := config.CreateTLSConfig(globalConfig.CACertFile, globalConfig.CertFile, globalConfig.KeyFile)
    conn, err := amqp.DialTLS(globalConfig.AmqpUrl, cfg)
    if err != nil {
        return err
    }
    defer conn.Close()

    channel, err := conn.Channel()
    if err != nil {
        return err
    }
    defer channel.Close()

    queue, err := channel.QueueDeclare("requests", false, false, false, false, nil)
    if err != nil {
        return err
    }

    err = channel.Qos(1, 0, false)
    if err != nil {
        return err
    }

    msgs, err := channel.Consume(queue.Name, "", false, false, false, false, nil)
    if err != nil {
        return err
    }

    forever := make(chan bool)
    go func() {
        for d := range msgs {
            log.Println("[CLIENT][AMQP] Received a message-request")
            responseBuffer, _ := proxyConn(d.Body)
            err = channel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
                ContentType:   "text/plain",
                CorrelationId: d.CorrelationId,
                Body:          responseBuffer.Bytes(),
            })
            if err != nil {
                log.Println("[CLIENT][AMQP] Failed to reply message-request: ", err.Error())
            }
            d.Ack(false)
        }
    }()

    log.Println("[CLIENT][AMQP] Waiting for messages. To exit press CTRL+C")
    <-forever
    return nil
}
