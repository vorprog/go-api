package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/richardsnider/golang/hello_world/util"
)

var appVersion = `1.0`
var processStartTime = time.Now()

var guidGeneratorURL = "http://www.uuidgenerator.net/api/version1"

type appMetaData struct {
	EnvironmentConfiguration string    `json:"EnvironmentConfiguration"`
	Version                  string    `json:"Version"`
	BuildGitCommit           string    `json:"BuildGitCommit"`
	BuildDateVersion         string    `json:"BuildDateVersion"`
	Hostname                 string    `json:"Hostname"`
	ProcessStartTime         time.Time `json:"ProcessStartTime"`
	CurrentTimestamp         time.Time `json:"CurrentTimestamp"`
	RequestGUID              string    `json:"RequestGUID"`
}

var hostname, hostnameError = os.Hostname()

var healthCheckMetaData = appMetaData{
	EnvironmentConfiguration: os.Getenv("APP_ENVIRONMENT_CONFIGURATION"),
	Version:                  appVersion,
	BuildGitCommit:           util.BuildCommitLinkerFlag,
	BuildDateVersion:         util.BuildDateVersionLinkerFlag,
	Hostname:                 hostname,
	ProcessStartTime:         processStartTime,
}

func rootPathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
}

func healthCheckHander(w http.ResponseWriter, r *http.Request) {
	healthCheckMetaData.CurrentTimestamp = time.Now()
	healthCheckMetaData.RequestGUID = util.GetURL(guidGeneratorURL)
	metaDataBytes, _ := json.MarshalIndent(healthCheckMetaData, "", "  ")
	fmt.Fprintln(w, string(metaDataBytes))
}

// Start starts the web server
func Start(int port) {
	if hostnameError != nil {
		log.Fatal(hostnameError)
	}

	http.HandleFunc("/", rootPathHandler)
	http.HandleFunc("/healthcheck", healthCheckHander)

	portSpecification := ":" + *port
	serverError := http.ListenAndServe(portSpecification, nil)
	if serverError != nil {
		log.Fatal(serverError)
	}
}
