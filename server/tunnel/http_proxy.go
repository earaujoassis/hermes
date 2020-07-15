package tunnel

import (
    "bytes"
    "fmt"
    "log"
    "net"

    "github.com/earaujoassis/hermes/tcp"
)

func handle(clientConn net.Conn) {
    defer clientConn.Close()
    log.Println(fmt.Sprintf("[TUNNEL] Received a connection from: %s", clientConn.RemoteAddr()))
    for {
        requestBuffer, err := tcp.ReadConn(clientConn)
        if requestBuffer.Len() > 0 {
            // fmt.Printf("%#v\n", requestBuffer.String())
            proxyConn(clientConn, requestBuffer)
        } else if err != nil {
            break
        }
    }
}

func proxyConn(clientConn net.Conn, requestBuffer bytes.Buffer) {
    responseBuffer, err := dispatchRequest(requestBuffer.Bytes())
    if err != nil {
        log.Println("[TUNNEL] Failed to proxy messages: ", err.Error())
        errorResponse(clientConn)
        return
    }
    if responseBuffer.Len() > 0 {
        // fmt.Printf("%#v\n", responseBuffer.String())
        _, err = clientConn.Write(responseBuffer.Bytes())
        if err != nil {
            log.Println("[TUNNEL] Failed to write response to client: ", err.Error())
            errorResponse(clientConn)
            return
        }
    } else {
        errorResponse(clientConn)
    }
}
