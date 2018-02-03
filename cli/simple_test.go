package main

import (
	"context"
	"github.com/thehivecorporation/log"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"
)

type server struct {
	index int
	wg    *sync.WaitGroup
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"host":        r.Host,
		"method":      r.Method,
		"remote-addr": r.RemoteAddr,
		"req-uri":     r.RequestURI,
		"url-host":    r.URL.Host,
		"url-path":    r.URL.Path,
		"scheme":      r.URL.Scheme,
	}).Info("Incoming request info")

	log.WithField("index", s.index).Info("Request received")

	s.wg.Done()
}

func Test_TCP_HTTP_MultiServerCall(t *testing.T) {
	log.Info("Launching test")

	go configureTargets("http://localhost:8080,http://localhost:80801")

	wg := &sync.WaitGroup{}
	wg.Add(2)

	s1 := &http.Server{Addr: ":8080", Handler: &server{index: 1, wg: wg}}
	s2 := &http.Server{Addr: ":8081", Handler: &server{index: 2, wg: wg}}

	go func() {
		log.Info("Starting server 1")
		if err := s1.ListenAndServe(); err != nil {
			log.WithError(err).Error("Error trying to close server 1")
			t.Fail()
		}
	}()
	go func() {
		log.Info("Starting server 2")

		if err := s2.ListenAndServe(); err != nil {
			log.WithError(err).Error("Error trying to close server 2")
			t.Fail()
		}
	}()

	addr := "http://localhost:8083"
	path := "/hello"
	params := "?key=value"

	go func() {
		time.Sleep(time.Second)

		if _, err := http.DefaultClient.Get(addr + path + params); err != nil {
			log.WithError(err).Error("Error received doing request")
			t.Fail()
		}
	}()

	wg.Wait()

	s1.Shutdown(context.Background())
	s2.Shutdown(context.Background())
}

func Test_UDP(t *testing.T) {
	conn, err := net.Dial("udp", "localhost:8080")
	if err != nil {
		log.WithError(err).Fatal("Opening tcp connection")
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Error(err)
		}
	}()

	if n, err := conn.Write([]byte("hello")); err != nil {
		log.WithError(err).Error("Error writing to target")
	} else {
		log.Debugf("%d bytes written", n)
	}
}
