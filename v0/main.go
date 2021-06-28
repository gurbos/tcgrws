package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	handler "github.com/gurbos/tcgrws/v0/handlers"
)

func main() {
	sigCh := make(chan os.Signal, 1)
	rtnCh := make(chan os.Signal, 1)
	sigChans := newSigChannels()
	sigChans.init(sigCh, rtnCh)

	signal.Notify(sigCh,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGHUP,
	)
	ctx := context.Background()
	c, cancel := context.WithCancel(ctx)
	go receiveSig(c, sigChans)

	r := mux.NewRouter()
	r.HandleFunc("/", handler.EndPointsHandler).Methods("GET")
	r.HandleFunc("/productLines", handler.ProductLineHandler).Methods("GET")
	r.HandleFunc("/metaData", handler.MetaDataHandler).Methods("GET")
	r.HandleFunc("/cards", handler.CardsHandler).Methods("GET")

	sh := newServHandler()
	sh.UseHandler(r)
	for {
		loadConfiguration()
		port := os.Getenv("PORT")
		if port == "" {
			port = "5000"
		}
		addr := ":" + port

		servConf := newServerConfig()
		servConf.init(sh, addr, 5*time.Second, 5*time.Second)
		server := startHttpServer(servConf)

		sig := <-rtnCh
		switch sig {
		case syscall.SIGHUP:
			if err := server.Shutdown(ctx); err != nil {
				panic(err)
			}
		case syscall.SIGTERM, syscall.SIGINT:
			cancel()
			if err := server.Shutdown(ctx); err != nil {
				panic(err)
			}
		}
	}
}
