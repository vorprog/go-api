package server

import (
	"net/http"
	"os"

	"github.com/richardsnider/golang/server/util"
)

func Start(port string) {
	util.Log("Server starting . . .")
	serverError := http.ListenAndServe(":"+port, http.HandlerFunc(baseHandler))
	util.Log("Listening on port " + port + " . . .")

	if serverError != nil {
		util.Log(serverError)
		os.Exit(1)
	}
}
