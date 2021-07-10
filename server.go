package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func startHttpServer(config *appConfigData) *http.Server {
	srv := &http.Server{
		Handler:      config.servHandler,
		Addr:         config.listenAddr + config.listenPort,
		ReadTimeout:  config.readTimeout,
		WriteTimeout: config.writeTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if _, ok := err.(*net.OpError); ok {
				log.Fatal("ListenAndServe():", err)
			}
			fmt.Println("ListenAndServe():", err)
		}
	}()
	return srv
}
