package util

import (
	"os"
	"time"
)

var ProcessStartTime = time.Now().UnixNano()
var AppVersion = `1.0`
var hostname, _ = os.Hostname()

const Http200Message = "HTTP 200 - OK"
const Http404Message = "HTTP 404 - Not Found"
const Http500Message = "HTTP 500 - Internal Server Error"

var BuildDateVersionLinkerFlag string
var BuildCommitLinkerFlag string

type AppMetaData struct {
	EnvironmentConfiguration string `json:"EnvironmentConfiguration"`
	Version                  string `json:"Version"`
	BuildGitCommit           string `json:"BuildGitCommit"`
	BuildDateVersion         string `json:"BuildDateVersion"`
	Hostname                 string `json:"Hostname"`
	ProcessStartTime         int64  `json:"ProcessStartTime"`
}

var CurrentAppMetaData = AppMetaData{
	EnvironmentConfiguration: os.Getenv("APP_ENVIRONMENT_CONFIGURATION"),
	Version:                  AppVersion,
	BuildGitCommit:           BuildCommitLinkerFlag,
	BuildDateVersion:         BuildDateVersionLinkerFlag,
	Hostname:                 hostname,
	ProcessStartTime:         ProcessStartTime,
}
