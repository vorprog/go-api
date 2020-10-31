package main

import (
	"encoding/json"
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
	Version          string    `json:"Version"`
	BuildGitCommit   string    `json:"BuildGitCommit"`
	BuildDateVersion string    `json:"BuildDateVersion"`
	Hostname         string    `json:"Hostname"`
	ProcessStartTime time.Time `json:"ProcessStartTime"`
	CurrentTimestamp time.Time `json:"CurrentTimestamp"`
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
		Version:          appVersion,
		BuildGitCommit:   buildCommitLinkerFlag,
		BuildDateVersion: buildDateVersionLinkerFlag,
		Hostname:         hostname,
		ProcessStartTime: processStartTime,
		CurrentTimestamp: time.Now(),
	}

	bytes, _ := json.MarshalIndent(healthCheckMetaData, "", "  ")
	fmt.Fprintln(w, string(bytes))
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
