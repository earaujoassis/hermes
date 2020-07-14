package tunnel

import (
    "fmt"
    "net"
)

func errorResponse(clientConn net.Conn) {
    body := `<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>502 Bad Gateway</title>
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
