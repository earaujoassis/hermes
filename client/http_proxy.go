package client

import (
    "bytes"
    "net"

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
