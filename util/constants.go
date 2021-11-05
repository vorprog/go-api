package util

import (
	"os"
	"time"
)

var ProcessStartTime = time.Now().UnixNano()
var AppVersion = "1.2"
var hostname, _ = os.Hostname()

const Http200Message = "HTTP 200 - OK"
const Http404Message = "HTTP 404 - Not Found"
const Http500Message = "HTTP 500 - Internal Server Error"

var BuildDateVersionLinkerFlag string
var BuildCommitLinkerFlag string

type AppMetaData struct {
	EnvironmentConfiguration string
	Version                  string
	BuildGitCommit           string
	BuildDateVersion         string
	Hostname                 string
	ProcessStartTime         int64
}

var CurrentAppMetaData = AppMetaData{
	EnvironmentConfiguration: os.Getenv("APP_ENVIRONMENT_CONFIGURATION"),
	Version:                  AppVersion,
	BuildGitCommit:           BuildCommitLinkerFlag,
	BuildDateVersion:         BuildDateVersionLinkerFlag,
	Hostname:                 hostname,
	ProcessStartTime:         ProcessStartTime,
}
