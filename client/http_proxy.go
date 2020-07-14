package client

import (
    "bytes"
    "net"
    "time"

    "github.com/earaujoassis/hermes/tcp"
    "github.com/earaujoassis/hermes/config"
)

func proxyConn(requestBuffer []byte) (bytes.Buffer, error) {
    var responseBuffer bytes.Buffer
    tunnelConn, err := net.Dial("tcp", config.LoadConfig().ClientHandlerServer)
    if err != nil {
        return responseBuffer, err
    }
    defer tunnelConn.Close()
    tunnelConn.SetReadDeadline(time.Now().Add(time.Millisecond * 200))

    _, err = tunnelConn.Write(requestBuffer)
    if err != nil {
        return responseBuffer, err
    }

    responseBuffer, err = tcp.ReadConn(tunnelConn)
    if err != nil {
        return responseBuffer, err
    }

    return responseBuffer, nil
}
