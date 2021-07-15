package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
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

// servHandler is middleware used to set the http.ResponseWriter headers.
// It implements the http.Hander interface and sets the response headers
// for every incoming http.Request.
type servHandler struct {
	baseHandler http.Handler
}

func (sh *servHandler) UseHandler(h http.Handler) {
	sh.baseHandler = h
}

func (sh *servHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dateStr := time.Now().Format(time.RFC1123) // For Date header
	sp := strings.Split(r.URL.Path[1:], "/")
	switch sp[0] {
	case "images":
		w.Header().Set("Content-Type", "image/jpg")
	default:
		w.Header().Set("Content-Type", "application/json")
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Date", dateStr)
	sh.baseHandler.ServeHTTP(w, r)
}

func newServHandler(bh http.Handler) http.Handler {
	sh := new(servHandler)
	sh.UseHandler(bh)
	return sh
}
