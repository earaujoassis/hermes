package tunnel

import (
    "log"
    "net"
)

func SetupTunnel() {
    listenerAddr, _ := net.ResolveTCPAddr("tcp", ":8080")
    listener, err := net.ListenTCP("tcp", listenerAddr)
    if err != nil {
        log.Fatalln("[TUNNEL] Setup for tunnel failed; panic: ", err.Error())
    }
    log.Println("[TUNNEL] Listening on port :8080")
    defer listener.Close()

    for {
        conn, err := listener.AcceptTCP()
        if err != nil {
            log.Println("[TUNNEL] Error while accepting connection: ", err.Error())
            continue
        }

        go handle(conn)
    }
}
