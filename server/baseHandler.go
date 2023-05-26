package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/vorprog/go-api/datastore"
	"github.com/vorprog/go-api/util"
)

func baseHandler(responseWriter http.ResponseWriter, request *http.Request) {
	requestStartTimestamp := time.Now().UTC().UnixNano()
	requestInfo := map[string]interface{}{
		"id ":              util.GetUuid(),
		"method":           request.Method,
		"path":             request.URL.Path,
		"requestIpAddress": request.RemoteAddr,
	}
	util.Log(requestInfo)
	datastore.Store("INSERT into events (id, method, path, request_ip_address) VALUES (?, ?, ?, ?)", requestInfo["id"], requestInfo["method"], requestInfo["path"], requestInfo["requestIpAddress"])

	responseWriter.Header().Add("Request-Id", requestInfo["id"].(string))
	var responseStatusCode int = 200
	var handlerResult interface{}

	if request.URL.Path == "/" || request.URL.Path == "/health" || request.URL.Path == "/healthcheck" {
		handlerResult = util.CurrentAppMetaData
	} else if request.URL.Path == "/bitcoin" {
		responseStatusCode, handlerResult = bitcoinHandler(requestInfo["id"].(string))
	} else if request.URL.Path == "/websocket" {
		serveWebsocket(responseWriter, request)
	} else {
		responseStatusCode = 404
		handlerResult = util.Http404Message
	}

	var responseContent []byte

	handlerResultString, handlerResultIsString := handlerResult.(string)
	if handlerResultIsString {
		responseContent = []byte(handlerResultString)
	} else {
		var jsonMarshalError error
		responseContent, jsonMarshalError = json.Marshal(handlerResult)

		if jsonMarshalError != nil {
			util.Log(jsonMarshalError)
			responseStatusCode = 500
			responseContent = []byte(util.Http500Message)
		}
	}

	responseWriter.WriteHeader(responseStatusCode)
	responseWriter.Write(responseContent)

	util.Log(map[string]interface{}{
		"requestId ":         requestInfo["id"].(string),
		"requestProcessTime": time.Now().UTC().UnixNano() - requestStartTimestamp,
		"responseStatusCode": responseStatusCode,
		"result":             string(responseContent),
	})
}
