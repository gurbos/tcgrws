package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gurbos/tcgrws/v0/dbio"
	res "github.com/gurbos/tcgrws/v0/resources"
)

func loadConfiguration() {
	var env environment
	env.loadEnvironment()

	res.ImagesHost = env.imagesHost
	res.ImagesDir = env.imagesDir
	dbio.DataSource.Init(
		env.dbHost, env.dbPort, env.dbUser, env.dbPasswd, env.dbName,
	)
	dbconn, err := dbio.DBConnection().DB()
	if err != nil {
		log.Fatal(err)
	}
	dbconn.SetMaxOpenConns(10)
	dbconn.SetConnMaxIdleTime(5)
	dbconn.SetConnMaxIdleTime(time.Hour)
}

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
