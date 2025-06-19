package client

import (
	"bytes"
	"net"

	"github.com/earaujoassis/hermes/config"
	"github.com/earaujoassis/hermes/tcp"
)

func proxyConn(requestBuffer []byte) (bytes.Buffer, error) {
	var responseBuffer bytes.Buffer
	var globalConfig config.Config = config.GetGlobalConfig()
	tunnelConn, err := net.Dial("tcp", globalConfig.ClientHandlerServer)
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
