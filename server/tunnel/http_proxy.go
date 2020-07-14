package tunnel

import (
    "bytes"
    "fmt"
    "log"
    "net"
    "time"
)

func handle(clientConn net.Conn) {
    defer clientConn.Close()
    log.Println(fmt.Sprintf("[TUNNEL] Received a connection from: %s", clientConn.RemoteAddr()))
    for {
        clientConn.SetReadDeadline(time.Now().Add(time.Millisecond * 200))
        requestBuffer, err := readConn(clientConn)
        if requestBuffer.Len() > 0 {
            // fmt.Printf("%#v\n", requestBuffer.String())
            proxyConn(clientConn, requestBuffer)
        } else if err != nil {
            break
        }
    }
}

func readConn(clientConn net.Conn) (bytes.Buffer, error) {
    requestBuffer := &bytes.Buffer{}
    for {
        data := make([]byte, 256)
        numBytes, err := clientConn.Read(data)
        if err != nil {
            return *requestBuffer, err
        }
        requestBuffer.Write(data[:numBytes])
    }
}

func proxyConn(clientConn net.Conn, requestBuffer bytes.Buffer) {
    tunnelConn, err := net.Dial("tcp", "localhost:8484")
    if err != nil {
        log.Println("[TUNNEL] Error requesting to remote endpoint: ", err.Error())
        errorResponse(clientConn)
        return
    }
    defer tunnelConn.Close()
    tunnelConn.SetReadDeadline(time.Now().Add(time.Millisecond * 200))

    _, err = tunnelConn.Write(requestBuffer.Bytes())
    if err != nil {
        log.Println("[TUNNEL] Error requesting to remote endpoint: ", err.Error())
        errorResponse(clientConn)
        return
    }

    responseBuffer, _ := readConn(tunnelConn)
    if responseBuffer.Len() > 0 {
        // fmt.Printf("%#v\n", responseBuffer.String())
        _, err = clientConn.Write(responseBuffer.Bytes())
        if err != nil {
            log.Println("[TUNNEL] Error writing response to client: ", err.Error())
            errorResponse(clientConn)
            return
        }
    }
}
