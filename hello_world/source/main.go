package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var gitCommitHash string
var processStartTime = time.Now()

type appMetaData struct {
	gitCommit        string
	hostname         string
	processID        int
	processStartTime time.Time
	currentTimestamp time.Time
}

func rootPathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
}

func healthCheckHander(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()

	healthCheckMetaData := appMetaData{
		gitCommit:        gitCommitHash,
		hostname:         hostname,
		processID:        os.Getpid(),
		processStartTime: processStartTime,
		currentTimestamp: time.Now(),
	}

	// TODO: fix json marshalling
	// bytes, _ := json.Marshal(healthCheckData)
	// fmt.Fprintln(w, string(bytes))

	fmt.Fprintln(w, "hostname: "+healthCheckMetaData.hostname)
	fmt.Fprintln(w, "git commit: "+healthCheckMetaData.gitCommit)
	fmt.Fprintln(w, "process id: "+strconv.Itoa(healthCheckMetaData.processID))
	fmt.Fprintln(w, "process start time: "+healthCheckMetaData.processStartTime.String())
	fmt.Fprintln(w, "request time: "+healthCheckMetaData.currentTimestamp.String())
}

func main() {
	http.HandleFunc("/", rootPathHandler)
	http.HandleFunc("/healthcheck", healthCheckHander)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal(err)
	}
}
