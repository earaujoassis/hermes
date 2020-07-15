package tcp

import (
    "bytes"
    "net"
    "time"
)

func ReadConn(clientConn net.Conn) (bytes.Buffer, error) {
    clientConn.SetReadDeadline(time.Now().Add(time.Millisecond * 200))
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
