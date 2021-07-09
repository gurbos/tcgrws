package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/gurbos/tcgrws/dbio"
	"github.com/gurbos/tcgrws/endpoints"
	"github.com/gurbos/tcgrws/handlers"
	handler "github.com/gurbos/tcgrws/handlers"
	"github.com/joho/godotenv"
)

type appConfigData struct {
	dbHost        string
	dbPort        string
	dbUser        string
	dbPass        string
	dbName        string
	host          string
	staticContent string
	listenAddr    string
	readTimeout   time.Duration
	writeTimeout  time.Duration
	handler       http.Handler
}

func (acd *appConfigData) loadConfigData() {
	var parent string
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if pwd := filepath.Base(wd); pwd == "bin" {
		parent = filepath.Dir(wd)
	} else {
		parent = wd
	}

	configFile := parent + "/config.env"
	envMap, err := godotenv.Read(configFile)
	if err != nil {
		log.Fatal(err)
	}

	acd.dbHost = envMap["DB_HOST"]
	acd.dbPort = envMap["DB_PORT"]
	acd.dbUser = envMap["DB_USER"]
	acd.dbPass = envMap["DB_PASS"]
	acd.dbName = envMap["DB_NAME"]
	acd.host = envMap["PUBLIC_HOSTNAME"]
	acd.staticContent = envMap["STATIC_CONTENT"]
	acd.listenAddr = envMap["LISTEN_ADDRESS"]
}

func (acd *appConfigData) configurePackages() {
	dbio.Configure(acd.dbHost, acd.dbPort, acd.dbUser, acd.dbPass, acd.dbName)
	endpoints.Configure(acd.host)
	handlers.Configure(acd.staticContent)
}

func (acd *appConfigData) loadAndConfigure() {
	acd.loadConfigData()
	acd.configurePackages()
}

func (acd *appConfigData) configureRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.EndPointsHandler).Methods("GET")
	r.HandleFunc("/productLines", handler.ProductLineHandler).Methods("GET")
	r.HandleFunc("/metaData", handler.MetaDataHandler).Methods("GET")
	r.HandleFunc("/cards", handler.CardsHandler).Methods("GET")
	handler := newServHandler()
	handler.UseHandler(r)
}

func newSigChannels() *sigChannels {
	return new(sigChannels)
}

type sigChannels struct {
	sigChan    <-chan os.Signal
	notifyChan chan os.Signal
}

func (sc *sigChannels) init(sch chan os.Signal, dch chan os.Signal) {
	sc.sigChan = sch
	sc.notifyChan = dch
}

func receiveSig(ctx context.Context, ch *sigChannels) {
	for {
		sig := <-ch.sigChan
		ch.notifyChan <- sig
		if val := <-ctx.Done(); val == struct{}{} {
			fmt.Println("receiveSig", ctx.Err().Error())
			return
		}
	}
}

func setupSignalHandling(sigCtx context.Context) chan os.Signal {
	sigCh := make(chan os.Signal, 1)
	notifyCh := make(chan os.Signal, 1)
	sigChans := newSigChannels()
	sigChans.init(sigCh, notifyCh)

	signal.Notify(
		sigCh,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGHUP,
	)
	go receiveSig(sigCtx, sigChans)
	return notifyCh
}
