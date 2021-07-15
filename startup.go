package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/gurbos/tcgrws/dbio"
	"github.com/gurbos/tcgrws/endpoints"
	"github.com/gurbos/tcgrws/handlers"
	"github.com/joho/godotenv"
)

type appConfigData struct {
	dbHost          string
	dbPort          string
	dbUser          string
	dbPass          string
	dbName          string
	maxOpenConns    int
	maxIdleConns    int
	maxConnLifetime time.Duration
	maxConnIdleTime time.Duration
	host            string
	staticContent   string
	listenAddr      string
	listenPort      string
	readTimeout     time.Duration
	writeTimeout    time.Duration
	servHandler     http.Handler
}

func (acd *appConfigData) loadConfigData() {
	configFile := os.Getenv("TCGRWS_CONFIG_FILE") // Get absolute path to configuration file
	envMap, err := godotenv.Read(configFile)
	if err != nil {
		panic("godotenv.Read(): " + err.Error())
	}

	openConns, err := strconv.Atoi(envMap["MAX_OPEN_CONNS"])
	idleConns, err := strconv.Atoi(envMap["MAX_IDLE_CONNS"])
	connLifetime, err := strconv.Atoi(envMap["MAX_CONN_LIFETIME_MIN"])
	connIdleTime, err := strconv.Atoi(envMap["MAX_CONN_IDLE_TIME_MIN"])
	rt, err := strconv.Atoi(envMap["READ_TIMEOUT_SEC"])
	wt, err := strconv.Atoi(envMap["WRITE_TIMEOUT_SEC"])
	if err != nil {
		panic("strconv.Atoi(): " + err.Error())
	}

	acd.readTimeout = time.Duration(rt * int(time.Second))
	acd.writeTimeout = time.Duration(wt * int(time.Second))
	acd.dbHost = envMap["DB_HOST"]
	acd.dbPort = envMap["DB_PORT"]
	acd.dbUser = envMap["DB_USER"]
	acd.dbPass = envMap["DB_PASS"]
	acd.dbName = envMap["DB_NAME"]
	acd.host = envMap["PUBLIC_HOSTNAME"]
	acd.staticContent = envMap["STATIC_CONTENT"]
	acd.listenAddr = envMap["LISTEN_ADDRESS"]
	acd.listenPort = envMap["LISTEN_PORT"]
	acd.maxOpenConns = openConns
	acd.maxIdleConns = idleConns
	acd.maxConnLifetime = time.Duration(connLifetime * int(time.Minute))
	acd.maxConnIdleTime = time.Duration(connIdleTime * int(time.Minute))
}

func (acd *appConfigData) configurePackages() {
	dbio.Configure(acd.dbHost, acd.dbPort, acd.dbUser,
		acd.dbPass, acd.dbName, acd.maxOpenConns,
		acd.maxIdleConns, acd.maxConnLifetime,
		acd.maxConnIdleTime,
	)
	handlers.Configure(acd.staticContent)
	endpoints.Configure(acd.host)
}

func (acd *appConfigData) configureRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.ApiRefHandler).Methods("GET")
	r.HandleFunc("/productLines", handlers.ProductLineHandler).Methods("GET")
	r.HandleFunc("/metaData", handlers.MetaDataHandler).Methods("GET")
	r.HandleFunc("/cards", handlers.CardsHandler).Methods("GET")
	r.HandleFunc("/images/{name}", handlers.ImagesHandler).Methods("GET")
	acd.servHandler = newServHandler(r)
}

func (acd *appConfigData) loadAndConfigure() {
	acd.loadConfigData()
	acd.configurePackages()
	acd.configureRoutes()
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

func sigHandler(ctx context.Context, ch *sigChannels) {
	for {
		sig := <-ch.sigChan
		ch.notifyChan <- sig
		if val := <-ctx.Done(); val == struct{}{} {
			fmt.Println("sigHandler", ctx.Err().Error())
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

	go sigHandler(sigCtx, sigChans)
	return notifyCh
}
