package main

import (
	"log"
	"os"
)

type environment struct {
	dbHost     string
	dbPort     string
	dbUser     string
	dbPasswd   string
	dbName     string
	imagesHost string
	imagesDir  string
}

func (e *environment) loadEnvironment() {
	var pres bool
	if e.dbHost = os.Getenv("TCG_DB_HOST"); e.dbHost == "" {
		log.Fatal("Error: 'TCG_DB_HOST' environment variable not set or does not exist!")
	}
	if e.dbPort = os.Getenv("TCG_DB_PORT"); e.dbPort == "" {
		log.Fatal("Error: 'TCG_DB_PORT' environment variable not set or does not exist!")
	}
	if e.dbUser = os.Getenv("TCG_DB_USER"); e.dbUser == "" {
		log.Fatal("Error: 'TCG_DB_USER' environment variable not set or does not exist!")
	}
	if e.dbPasswd = os.Getenv("TCG_DB_PASSWD"); e.dbPasswd == "" {
		log.Fatal("Error: 'TCG_DB_PASSWD' environment variable not set or does not exist!")
	}
	if e.dbName = os.Getenv("TCG_DB_NAME"); e.dbName == "" {
		log.Fatal("Error: 'TCG_DB_NAME' environment variable not set or does not exist!")
	}
	e.imagesHost, pres = os.LookupEnv("TCG_IMAGES_HOST")
	if !pres {
		log.Fatal("TCG_IMAGES_HOST environment variable not present")
	}
	e.imagesDir, pres = os.LookupEnv("TCG_IMAGES_DIR")
	if !pres {
		log.Fatal("TCG_IMAGES_DIR environment variable not present")
	}
}
