package main

import (
	"github.com/thehivecorporation/log"
	"github.com/thehivecorporation/log/writers/json"
	"os"
	"strings"
)

func launch() {
	log.SetLevel(log.LevelStrings[appConfig.LogLevel])
	if appConfig.LogOutput == "json" {
		log.SetWriter(json.New(os.Stderr))
	}

	ts := strings.Split(appConfig.TargetsString, ",")
	if len(ts) == 0 {
		log.Fatal("Not enough targets. Use '--targets' to specify at least one")
	}

	if appConfig.Mode == "udp" {
		udpServer(ts)
	} else if appConfig.Mode == "tcp" {
		tcpServer(ts)
	} else {
		log.Fatalf("'%s' mode not recognized. Use one of 'udp' or 'tcp'")
	}
}