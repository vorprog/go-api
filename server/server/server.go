package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/richardsnider/golang/server/util"
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
	util.Log()
	fmt.Fprintln(w, "hello world!")
}

func healthCheckHander(w http.ResponseWriter, r *http.Request) {
	healthCheckMetaData.CurrentTimestamp = time.Now()
	healthCheckMetaData.RequestGUID, getUrlError = util.GetURL(guidGeneratorURL)

	if getUrlError {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	metaDataBytes, _ := json.Marshal(healthCheckMetaData)
	metaData := string(metaDataBytes)
	util.Log(metaData)
	fmt.Fprintln(w, metaData)
}

func Start(port *string) {
	if hostnameError != nil {
		util.LogError(hostnameError)
		os.Exit(1)
	}

	http.HandleFunc("/", rootPathHandler)
	http.HandleFunc("/healthcheck", healthCheckHander)

	portSpecification := ":" + *port
	util.Log("port is " + *port)

	serverError := http.ListenAndServe(portSpecification, nil)
	if serverError != nil {
		util.LogError(serverError)
		os.Exit(1)
	}
}
