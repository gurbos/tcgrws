package main

import (
	"net/http"
	"time"
)

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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	{
		dateStr := time.Now().Format(time.RFC1123)
		w.Header().Set("Date", dateStr)
	}
	sh.baseHandler.ServeHTTP(w, r)
}

func newServHandler(bh http.Handler) http.Handler {
	sh := new(servHandler)
	sh.UseHandler(bh)
	return sh
}
