package server

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	util.Log(map[string]interface{}{
		"requestId ":       requestId,
		"method":           request.Method,
		"path":             request.URL.Path,
		"requestIpAddress": request.RemoteAddr,
	})

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
	util.Log(map[string]interface{}{
		"requestId ":  result.RequestId,
		"requestTime": fmt.Sprint(requestProcessTime),
		"result":      string(responseContent),
	})
}
