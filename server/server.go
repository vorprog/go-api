package server

import (
	"net/http"
	"reflect"

	_ "net/http/pprof"

	"github.com/vorprog/go-api/datastore"
	"github.com/vorprog/go-api/util"
)

func Start(port string) error {
	err := datastore.Init(map[string]reflect.Type{}) // TODO: pass in dataset types
	if err != nil {
		return err
	}

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
