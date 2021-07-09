package main

import (
	"log"
	"net/http"
)

func startHttpServer(config *appConfigData) *http.Server {
	srv := &http.Server{
		Handler:      config.handler,
		Addr:         config.listenAddr,
		ReadTimeout:  config.readTimeout,
		WriteTimeout: config.writeTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("ListenAndServe(): ", err)
		}
	}()
	return srv
}
