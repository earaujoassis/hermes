package tunnel

import (
    "bufio"
    "bytes"
    "fmt"
    "log"
    "net"
    "strings"
)

func handle(clientConn *net.TCPConn) {
    defer clientConn.Close()
    log.Println(fmt.Sprintf("[TUNNEL] Received a connection from: %s", clientConn.RemoteAddr()))
    requestData := request(clientConn)
    response(clientConn, requestData)
}

func request(clientConn *net.TCPConn) string {
    var requestBuffer bytes.Buffer
    scanner := bufio.NewScanner(clientConn)
    log.Println(fmt.Sprintf("[TUNNEL] Building request for client: %s", clientConn.RemoteAddr()))
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "Host:") {
            requestBuffer.WriteString("Host: localhost:9000\r\n")
        } else {
            requestBuffer.WriteString(strings.TrimSpace(line) + "\r\n")
        }
        // Reached end of scanning/header
        if line == "" {
            break
        }
    }

    log.Println(fmt.Sprintf("[TUNNEL] Request is complete for client: %s", clientConn.RemoteAddr()))
    return requestBuffer.String()
}

func response(clientConn *net.TCPConn, requestData string) {
    var buildingHeader bool

    tunnelAddr, _ := net.ResolveTCPAddr("tcp", "localhost:9000")
    tunnelConn, err := net.DialTCP("tcp", nil, tunnelAddr)
    if err != nil {
        log.Println("[TUNNEL] Error while creating connection to remote endpoint: ", err.Error())
        errorResponse(clientConn)
        return
    }
    defer tunnelConn.Close()

    log.Println(fmt.Sprintf("[TUNNEL] Sending request to remote endpoint for client: %s", clientConn.RemoteAddr()))
    _, err = tunnelConn.Write([]byte(requestData))
    if err != nil {
        log.Println("[TUNNEL] Error writing request to remote endpoint: ", err.Error())
        errorResponse(clientConn)
        return
    }

    log.Println(fmt.Sprintf("[TUNNEL] Piping response from remote endpoint to client: %s", clientConn.RemoteAddr()))
    scanner := bufio.NewScanner(tunnelConn)
    buildingHeader = true
    for scanner.Scan() {
        line := scanner.Text()
        if buildingHeader && line != "" {
            if _, err = clientConn.Write([]byte(line + "\r\n")); err != nil {
                log.Println("[TUNNEL] Error writing response to client: ", err.Error())
            }
        } else if line == "" && buildingHeader {
            if _, err = clientConn.Write([]byte("\r\n")); err != nil {
                log.Println("[TUNNEL] Error writing response to client: ", err.Error())
            }
            buildingHeader = false
        } else {
            if _, err = clientConn.Write([]byte(line + "\n")); err != nil {
                log.Println("[TUNNEL] Error writing response to client: ", err.Error())
            }
        }
    }

    clientConn.CloseWrite()
}

func errorResponse(clientConn net.Conn) {
    body := `
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>502 Bad Gatway</title>
    </head>
    <body>
        <div style="margin:20px auto;max-width:90%">
            <h1 style="text-align:center;padding-bottom:20px;border-bottom:1px solid black">502 Bad Gateway</h1>
        </div>
    </body>
</html>
`
    fmt.Fprint(clientConn, "HTTP/1.1 502 Bad Gateway\r\n")
    fmt.Fprintf(clientConn, "Content-Length: %d\r\n", len(body))
    fmt.Fprint(clientConn, "Content-Type: text/html\r\n")
    fmt.Fprint(clientConn, "\r\n")
    fmt.Fprint(clientConn, body)
}
