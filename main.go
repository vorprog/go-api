package main

import (
	"os"

	"github.com/vorprog/go-api/server"
	"github.com/vorprog/go-api/util"
)

func main() {
	util.Log("Loaded main module.")
	if os.Getenv("SOPS_FILE_URL") != "" {
		var sopsError = util.SetEnvironmentFromSopsURL()
		if sopsError != nil {
			util.Log(sopsError)
			os.Exit(1)
		}
	}

	awsIdentity := util.GetAwsIdentity()
	util.Log(awsIdentity)

	go util.Monitor()

	port := os.Getenv("APP_SERVER_PORT")
	server.Start(port)
}
