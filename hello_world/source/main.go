package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var appVersion = `1.0`
var buildDateVersionLinkerFlag string
var buildCommitLinkerFlag string
var processStartTime = time.Now()

type appMetaData struct {
	version          string
	buildGitCommit   string
	buildDateVersion string
	hostname         string
	processStartTime time.Time
	currentTimestamp time.Time
}

func rootPathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
}

func healthCheckHander(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	healthCheckMetaData := appMetaData{
		version:          appVersion,
		buildGitCommit:   buildCommitLinkerFlag,
		buildDateVersion: buildDateVersionLinkerFlag,
		hostname:         hostname,
		processStartTime: processStartTime,
		currentTimestamp: time.Now(),
	}

	// TODO: fix json marshalling
	// bytes, _ := json.Marshal(healthCheckData)
	// fmt.Fprintln(w, string(bytes))

	fmt.Fprintln(w, "version: "+healthCheckMetaData.version)
	fmt.Fprintln(w, "git commit: "+healthCheckMetaData.buildGitCommit)
	fmt.Fprintln(w, "build date: "+healthCheckMetaData.buildDateVersion)
	fmt.Fprintln(w, "hostname: "+healthCheckMetaData.hostname)
	fmt.Fprintln(w, "process start time: "+healthCheckMetaData.processStartTime.String())
	fmt.Fprintln(w, "request time: "+healthCheckMetaData.currentTimestamp.String())
}

func main() {
	http.HandleFunc("/", rootPathHandler)
	http.HandleFunc("/healthcheck", healthCheckHander)

	port := ":" + os.Getenv("PORT")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
