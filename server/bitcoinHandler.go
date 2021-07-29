package server

import (
	"github.com/vorprog/go-api/util"
)

func bitcoinHandler(requestId string) (responseStatusCode int, responseContent interface{}) {
	coindeskResponse, getUrlError := util.GetURL("https://api.coindesk.com/v1/bpi/historical/close.json")

	if getUrlError != nil {
		util.Log(getUrlError)
		return 500, "HTTP STATUS 500 - Internal Server Error"
	}

	return 200, coindeskResponse
}
