package tcp

import (
    "bytes"
    "net"
)

func ReadConn(clientConn net.Conn) (bytes.Buffer, error) {
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
