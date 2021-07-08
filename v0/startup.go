package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

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
}

func (acd *appConfigData) loadConfiguration() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(wd)
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

}

var ServerConfig *appConfigData

func newSigChannels() *sigChannels {
	return new(sigChannels)
}

type sigChannels struct {
	sigChan <-chan os.Signal
	rtnChan chan os.Signal
}

func (sc *sigChannels) init(sch chan os.Signal, dch chan os.Signal) {
	sc.sigChan = sch
	sc.rtnChan = dch
}

func receiveSig(ctx context.Context, ch *sigChannels) {
	for {
		sig := <-ch.sigChan
		ch.rtnChan <- sig
		if val := <-ctx.Done(); val == struct{}{} {
			fmt.Println("receiveSig", ctx.Err().Error())
			return
		}
	}
}
