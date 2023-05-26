package main

import (
	"strconv"
	"syscall"

	"github.com/vorprog/go-api/datastore"
	"github.com/vorprog/go-api/server"
	"github.com/vorprog/go-api/util"
)

func main() {
	pid, _, _ := syscall.Syscall(syscall.SYS_GETPID, 0, 0, 0)
	util.Log("process id: " + strconv.Itoa(int(pid)))
	go util.Monitor()
	util.InitConfig()

	_, err := datastore.Init()
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
