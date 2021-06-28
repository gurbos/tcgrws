package main

import (
	"os"
	"time"

	"github.com/gorilla/mux"
	handler "github.com/gurbos/tcgrws/v0/handlers"
)

func main() {
	startUp()

	r := mux.NewRouter()
	r.HandleFunc("/", handler.EndPointsHandler).Methods("GET")
	r.HandleFunc("/productLines", handler.ProductLineHandler).Methods("GET")
	r.HandleFunc("/metaData", handler.MetaDataHandler).Methods("GET")
	r.HandleFunc("/cards", handler.CardsHandler).Methods("GET")

	sh := newServHandler()
	sh.UseHandler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	addr := ":" + port

	servConf := newServerConfig()
	servConf.init(sh, addr, 5*time.Second, 5*time.Second)

	server := startHttpServer(servConf)
}
