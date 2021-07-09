package main

import (
	"net/http"
	"time"
)

// servHandler is middleware used to set the http.ResponseWriter headers.
// It implements the http.Hander interface and sets the response headers
// for every incoming http.Request.
type servHandler struct {
	handler http.Handler
}

func newServHandler() *servHandler {
	return new(servHandler)
}

func (sh *servHandler) UseHandler(h http.Handler) {
	sh.handler = h
}

func (sh *servHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	{
		dateStr := time.Now().Format(time.RFC1123)
		w.Header().Set("Date", dateStr)
	}
	sh.handler.ServeHTTP(w, r)
}
