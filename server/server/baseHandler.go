package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/richardsnider/golang/server/util"
)

type responseResult struct {
	RequestId             string
	RequestStartTimestamp int64
	HandlerResult         interface{}
}

func baseHandler(responseWriter http.ResponseWriter, request *http.Request) {
	requestStartTimestamp := time.Now().UTC().UnixNano()
	requestId := util.GetUuid()
	util.Log("Request ID " + requestId + " " + request.Method + " request to " + request.URL.Path + " from " + request.RemoteAddr)

	var responseStatusCode int = 200
	var handlerResult interface{}

	if request.URL.Path == "/" || request.URL.Path == "/health" || request.URL.Path == "/healthcheck" {
		handlerResult = util.CurrentAppMetaData
	} else if request.URL.Path == "/bitcoin" {
		responseStatusCode, handlerResult = bitcoinHandler(requestId)
	}

	result := responseResult{
		RequestId:             requestId,
		RequestStartTimestamp: requestStartTimestamp,
		HandlerResult:         handlerResult,
	}

	responseWriter.WriteHeader(responseStatusCode)
	responseContent, _ := json.Marshal(result)
	responseWriter.Write(responseContent)

	requestProcessTime := time.Now().UTC().UnixNano() - result.RequestStartTimestamp
	util.Log("Request ID " + result.RequestId + " took " + strconv.Itoa(int(requestProcessTime)) + "ns to respond with " + string(responseContent))
}
