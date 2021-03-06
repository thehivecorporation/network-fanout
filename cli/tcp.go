package main

import (
	"github.com/thehivecorporation/log"
	"net"
	"fmt"
)

func tcpServer(targets []string) {
	l, err := net.Listen(appConfig.Mode, fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port))
	if err != nil {
		log.WithError(err).Fatalf("Opening %s server", appConfig.Mode)
	}
	defer closeConn(l)

	for {
		c, err := l.Accept()
		if err != nil {
			log.WithError(err).Fatal("Error accepting connection")
			continue
		}

		go handleConnection(c, targets)
	}
}

func handleConnection(c net.Conn, ts []string) {
	defer closeConn(c)

	buffer := make([]byte, appConfig.ReadBufferSize)
	_, err := c.Read(buffer)
	if err != nil {
		log.WithError(err).Error("Error reading data")
	}

	for _, target := range ts {
		go handleTcp(target, buffer)
	}
}

func handleTcp(t string, byt []byte) {
	conn, err := net.Dial("tcp", t)
	if err != nil {
		log.WithError(err).Error("Error opening TCP connection")
		return
	}
	defer closeConn(conn)

	writeToConn(conn, byt)
}
