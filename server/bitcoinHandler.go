package server

import (
	"github.com/vorprog/go-api/util"
)

func bitcoinHandler(requestId string) (responseStatusCode int, responseContent interface{}) {
	coindeskResponse, getUrlError := util.GetURL("https://api.coindesk.com/v1/bpi/historical/close.json")

	if getUrlError != nil {
		util.Log(getUrlError)
		return 500, util.Http500Message
	}

	return 200, coindeskResponse
}
