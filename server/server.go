package server

import (
	"net/http"

	_ "net/http/pprof"

	"github.com/vorprog/go-api/util"
)

func Start(port string) error {

	util.Log(map[string]interface{}{
		"messsage ": "server starting",
		"port":      port,
	})

	serverError := http.ListenAndServe(":"+port, http.HandlerFunc(baseHandler))

	if serverError != nil {
		return serverError
	}

	return nil
}
