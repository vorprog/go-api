package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/vorprog/go-api/util"
)

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
	} else {
		responseStatusCode, handlerResult = NotFound()
	}

	responseWriter.Header().Add(`request-id`, requestId)
	responseWriter.WriteHeader(responseStatusCode)
	responseContent, _ := json.Marshal(handlerResult)
	responseWriter.Write(responseContent)

	requestProcessTime := time.Now().UTC().UnixNano() - requestStartTimestamp
	util.Log(map[string]interface{}{
		"requestId ":  requestId,
		"requestTime": fmt.Sprint(requestProcessTime),
		"result":      string(responseContent),
	})
}
