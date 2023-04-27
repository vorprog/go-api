package main

import (
	"reflect"

	"github.com/vorprog/go-api/datastore"
	"github.com/vorprog/go-api/server"
	"github.com/vorprog/go-api/util"
)

func main() {
	util.Log("Loaded main module.")
	go util.Monitor()
	util.InitConfig()

	err := datastore.Init(map[string]reflect.Type{})
	if err != nil {
		util.Log(err)
		panic(err)
	}

	awsIdentity, err := util.GetAwsIdentity()
	if err != nil {
		util.Log(err)
		panic(err)
	} else {
		util.Log(awsIdentity)
	}

	err = server.Start(util.Config.ServerPort)
	if err != nil {
		util.Log(err)
		panic(err)
	}
}
