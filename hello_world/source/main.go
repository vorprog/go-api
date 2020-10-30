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

type healthCheckData struct {
	gitCommit        string
	hostname         string
	processID        int
	processStartTime time.Time
	requestTime      time.Time
}

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()

	healthCheckData := healthCheckData{
		gitCommit:        gitCommitHash,
		hostname:         hostname,
		processID:        os.Getpid(),
		processStartTime: processStartTime,
		requestTime:      time.Now(),
	}

	// TODO: fix json marshalling
	// bytes, _ := json.Marshal(healthCheckData)
	// fmt.Fprintln(w, string(bytes))
	fmt.Fprintln(w, "hostname: "+healthCheckData.hostname)
	fmt.Fprintln(w, "git commit: "+healthCheckData.gitCommit)
	fmt.Fprintln(w, "process id: "+strconv.Itoa(healthCheckData.processID))
	fmt.Fprintln(w, "process start time: "+healthCheckData.processStartTime.String())
	fmt.Fprintln(w, "request time: "+healthCheckData.requestTime.String())
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
