package main

import (
	"log"
	"net/http"
	"time"
)

func newServerConfig() *serverConfig {
	return new(serverConfig)
}

type serverConfig struct {
	handler      http.Handler
	addr         string
	readTimeout  time.Duration
	writeTimeout time.Duration
}

func (sc *serverConfig) init(
	h http.Handler, a string, rt time.Duration, wt time.Duration) {
	sc.handler = h
	sc.addr = a
	sc.readTimeout = rt
	sc.writeTimeout = wt
}

func startHttpServer(config *serverConfig) *http.Server {
	srv := &http.Server{
		Handler:      config.handler,
		Addr:         config.addr,
		ReadTimeout:  config.readTimeout,
		WriteTimeout: config.writeTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("ListenAndServe() %v", err)
		}
	}()
	return srv
}
