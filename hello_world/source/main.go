package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var compiledGitCommitHash string
var compiledGitCommitTimestamp string
var processStartTime = time.Now()

type appMetaData struct {
	gitCommitHash      string
	gitCommitTimestamp string
	hostname           string
	processID          int
	processStartTime   time.Time
	currentTimestamp   time.Time
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
		gitCommitHash:      compiledGitCommitHash,
		gitCommitTimestamp: compiledGitCommitTimestamp,
		hostname:           hostname,
		processID:          os.Getpid(),
		processStartTime:   processStartTime,
		currentTimestamp:   time.Now(),
	}

	// TODO: fix json marshalling
	// bytes, _ := json.Marshal(healthCheckData)
	// fmt.Fprintln(w, string(bytes))

	fmt.Fprintln(w, "hostname: "+healthCheckMetaData.hostname)
	fmt.Fprintln(w, "git commit: "+healthCheckMetaData.gitCommitHash)
	fmt.Fprintln(w, "process id: "+strconv.Itoa(healthCheckMetaData.processID))
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
