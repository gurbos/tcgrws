package main

import (
	"fmt"
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
			fmt.Println("ListenAndServe(): ", err)
		}
	}()
	return srv
}
