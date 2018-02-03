package main

import (
	"net"
	"github.com/thehivecorporation/log"
	"io"
)

func closeConn(c io.Closer){
	if err := c.Close(); err != nil {
		log.Error(err)
	}
}

func writeToConn(c net.Conn, byt []byte){
	if n, err := c.Write(byt); err != nil {
		log.WithError(err).Error("Error writing to target")
	} else {
		log.Debugf("%d bytes written", n)
	}
}