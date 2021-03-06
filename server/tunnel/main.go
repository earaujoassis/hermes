package tunnel

import (
    "fmt"
    "log"
    "net"

    "github.com/earaujoassis/hermes/config"
)

func SetupTunnel() {
    listener, err := net.Listen("tcp", fmt.Sprintf(":%v", config.GetEnvVarDefault("PORT", "8080")))
    if err != nil {
        log.Fatalln("[TUNNEL] Panic: failed to initiate listener: ", err.Error())
    }
    log.Println(fmt.Sprintf("[TUNNEL] Listening on port :%v", config.GetEnvVarDefault("PORT", "8080")))
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println("[TUNNEL] Failed to accept connection: ", err.Error())
            continue
        }

        go handle(conn)
    }
}
