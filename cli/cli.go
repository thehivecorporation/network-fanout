package main

import (
	"context"
	"github.com/timjchin/unpuzzled"
	"net"
	"net/http"
	"os"
	"time"
)

var appConfig struct {
	Mode           string
	TargetsString  string
	ReadBufferSize int
	LogLevel       string
	LogOutput      string
	Port           int
}

var client = http.Client{
	Transport: &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			conn, err := (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial(network, addr)
			if err != nil {
				println("Error during DIAL:", err.Error())
			}
			return conn, err
		},
		TLSHandshakeTimeout: 10 * time.Second,
	},
	Timeout: 5 * time.Second,
}

func main() {
	app := unpuzzled.NewApp()
	app.Command = &unpuzzled.Command{
		Name: "server",
		Variables: []unpuzzled.Variable{
			&unpuzzled.StringVariable{Name: "targets", Destination: &appConfig.TargetsString, Required: true},
			&unpuzzled.StringVariable{Name: "log-output", Destination: &appConfig.LogOutput, Default: "text", Description: "Choose between text and json"},
			&unpuzzled.StringVariable{Name: "log", Destination: &appConfig.LogLevel, Default: "info",
				Description: "You can set from more logs to less logs: debug, info, warn, error or fatal"},
			&unpuzzled.StringVariable{Name: "mode", Destination: &appConfig.Mode, Default: "tcp", Description: "You can use tcp or udp"},
			&unpuzzled.IntVariable{Name: "port", Default: 8083, Destination: &appConfig.Port},
			&unpuzzled.IntVariable{Name: "read-buffer-size", Destination: &appConfig.ReadBufferSize, Default: 1024},
		},
		Action: launch,
	}

	app.Run(os.Args)
}
