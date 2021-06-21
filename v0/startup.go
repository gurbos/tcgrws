package main

import (
	"log"
	"time"

	"github.com/gurbos/tcgrws/v0/dbio"
	res "github.com/gurbos/tcgrws/v0/resources"
)

func startUp() {
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
