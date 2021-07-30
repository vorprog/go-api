package main

import (
	"os"

	"github.com/vorprog/go-api/server"
	"github.com/vorprog/go-api/util"
)

func main() {
	util.Log("Loaded main module.")

	awsIdentity := util.GetAwsIdentity()
	util.Log(awsIdentity)

	go util.Monitor()

	port := os.Getenv("APP_SERVER_PORT")
	server.Start(port)
}
