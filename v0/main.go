package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	handler "github.com/gurbos/tcgrws/v0/handlers"
)

// servHandler adds functionality to the handler reference by the 'mr' field.
// It implements the http.Handler interface. The 'ServeHTTP' method sets the
// HTTP header of the 'http.ResponseWriter' then calls the 'ServeHTTP' method
// that corresponds to the 'mr' field.
type servHandler struct {
	handler http.Handler
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

func main() {
	startUp()

	r := mux.NewRouter()
	r.HandleFunc("/", handler.APIHandler).Methods("GET")
	r.HandleFunc("/productLines", handler.ProductLineHandler).Methods("GET")
	r.HandleFunc("/metaData", handler.MetaDataHandler).Methods("GET")
	r.HandleFunc("/cards", handler.CardsHandler).Methods("GET")

	sh := servHandler{handler: r}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	server := http.Server{
		Handler:      &sh,
		Addr:         ":" + port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
