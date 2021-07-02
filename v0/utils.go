package main

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

type environment struct {
	dbHost      string
	dbPort      string
	dbUser      string
	dbPasswd    string
	dbName      string
	serviceHost string
	imagesHost  string
	imagesDir   string
}

func (e *environment) loadEnvironment() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(wd)
	configFile := parent + "/config.env"
	envMap, _ := godotenv.Unmarshal(configFile)

	// var pres bool
	// if e.dbHost = os.Getenv("TCG_DB_HOST"); e.dbHost == "" {
	// 	log.Fatal("Error: 'TCG_DB_HOST' environment variable not set or does not exist!")
	// }
	// if e.dbPort = os.Getenv("TCG_DB_PORT"); e.dbPort == "" {
	// 	log.Fatal("Error: 'TCG_DB_PORT' environment variable not set or does not exist!")
	// }
	// if e.dbUser = os.Getenv("TCG_DB_USER"); e.dbUser == "" {
	// 	log.Fatal("Error: 'TCG_DB_USER' environment variable not set or does not exist!")
	// }
	// if e.dbPasswd = os.Getenv("TCG_DB_PASSWD"); e.dbPasswd == "" {
	// 	log.Fatal("Error: 'TCG_DB_PASSWD' environment variable not set or does not exist!")
	// }
	// if e.dbName = os.Getenv("TCG_DB_NAME"); e.dbName == "" {
	// 	log.Fatal("Error: 'TCG_DB_NAME' environment variable not set or does not exist!")
	// }
	// e.imagesHost, pres = os.LookupEnv("TCG_IMAGES_HOST")
	// if !pres {
	// 	log.Fatal("TCG_IMAGES_HOST environment variable not present")
	// }
	// e.imagesDir, pres = os.LookupEnv("TCG_IMAGES_DIR")
	// if !pres {
	// 	log.Fatal("TCG_IMAGES_DIR environment variable not present")
	// }
}

// servHandler is middleware used to set the http.ResponseWriter headers.
// It implements the http.Hander interface and sets the response headers
// for every incoming http.Request.
type servHandler struct {
	handler http.Handler
}

func newServHandler() *servHandler {
	return new(servHandler)
}

func (sh *servHandler) UseHandler(h http.Handler) {
	sh.handler = h
}

func (sh *servHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	{
		dateStr := time.Now().Format(time.RFC1123)
		w.Header().Set("Date", dateStr)
	}
	sh.handler.ServeHTTP(w, r)
}
