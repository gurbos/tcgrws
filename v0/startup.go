package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type appConfigData struct {
	dbHost     string
	dbPort     string
	dbUser     string
	dbPasswd   string
	dbName     string
	imagesHost string
	imagesDir  string
}

func (acd *appConfigData) loadConfiguration() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(wd)
	configFile := parent + "/config.env"
	envMap, _ := godotenv.Unmarshal(configFile)

	acd.dbHost = envMap["DB_HOST"]
	acd.dbPort = envMap["DB_PORT"]
	acd.dbUser = envMap["DB_USER"]
	acd.dbPasswd = envMap["DB_PASSWD"]
	acd.dbName = envMap["DB_NAME"]
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
