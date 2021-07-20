package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/richardsnider/golang/server/util"
)

type requestMetaData struct {
	AppMetaData      util.AppMetaData `json:"AppMetaData"`
	CurrentTimestamp time.Time        `json:"CurrentTimestamp"`
	RequestGUID      string           `json:"RequestGUID"`
}

var guidRegexp *regexp.Regexp = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

func healthCheckHander(w http.ResponseWriter, r *http.Request) {
	util.Log("Request from " + r.RemoteAddr)
	timeOfRequest := time.Now()

	requestGuid, getUrlError := util.GetURL("http://www.uuidgenerator.net/api/version1")

	if getUrlError != nil {
		util.Log(getUrlError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP STATUS 500 - Internal Server Error"))
		return
	}

	if !guidRegexp.MatchString(requestGuid) {
		util.Log("GUID request resulted in " + requestGuid + ", which is not a valid GUID")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP STATUS 500 - Internal Server Error"))
		return
	}

	healthCheckMetaData := requestMetaData{
		AppMetaData:      util.CurrentAppMetaData,
		CurrentTimestamp: timeOfRequest,
		RequestGUID:      requestGuid,
	}

	metaDataBytes, _ := json.Marshal(healthCheckMetaData)
	response := string(metaDataBytes)
	fmt.Fprintln(w, response)
	util.Log("Response " + response)
}
