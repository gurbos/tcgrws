package main

import (
	"log"
	"os"
	"strconv"
)

type environment struct {
	dbHost   string
	dbPort   string
	dbUser   string
	dbPasswd string
	dbName   string
}

func (e *environment) loadEnvironment() {
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
}

func joinInt64ListAsStr(list []int64, sep string) string {
	numStr := strconv.Itoa(int(list[0]))
	for _, elem := range list[1:] {
		numStr += sep + strconv.Itoa(int(elem))
	}

	return numStr
}
