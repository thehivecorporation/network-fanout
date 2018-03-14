package main

import (
	"fmt"
	"github.com/thehivecorporation/log"
	"net"
	"strconv"
	"strings"
)

func udpServer(ts []string) {
	pc, err := net.ListenPacket("udp", fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := pc.Close(); err != nil {
			log.Error(err)
		}
	}()

	targets := parseUdpAddresses(ts)
	if targets == nil {
		return
	}

	for {
		buffer := make([]byte, appConfig.ReadBufferSize)
		n, _, err := pc.ReadFrom(buffer)
		if err != nil {
			log.WithError(err).Fatal("Could not read UDP package")
			continue
		}

		for _, target := range targets {
			go handleUdp(target, buffer[:n])
		}
	}
}

func parseUdpAddresses(ts []string) []net.UDPAddr {
	targets := make([]net.UDPAddr, len(ts))
	for i, t := range ts {
		hostPort := strings.Split(t, ":")

		port, err := strconv.Atoi(hostPort[1])
		if err != nil {
			log.WithError(err).Fatal("Could not parse port")
			return nil
		}

		targets[i] = net.UDPAddr{IP: net.ParseIP(hostPort[0]), Port: port}
	}

	return targets
}

func handleUdp(t net.UDPAddr, byt []byte) {
	conn, err := net.DialUDP("udp", nil, &t)
	if err != nil {
		log.WithError(err).Fatal("Error opening UDP connection")
		return
	}
	defer closeConn(conn)

	writeToConn(conn, byt)
}
