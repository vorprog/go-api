package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/richardsnider/golang/server/util"
)

func rootPathHandler(w http.ResponseWriter, r *http.Request) {
	util.Log("Request from " + r.RemoteAddr)
	fmt.Fprintln(w, "hello world!")
}

func Start(port string) {
	util.Log("Server starting . . .")
	http.HandleFunc("/", rootPathHandler)
	http.HandleFunc("/healthcheck", healthCheckHander)
	serverError := http.ListenAndServe(":"+port, nil)
	util.Log("Listening on port " + port + " . . .")

	if serverError != nil {
		util.Log(serverError)
		os.Exit(1)
	}
}
