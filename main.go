package main

import (
	"log"
	"os"

	"github.com/samber/lo"
	"github.com/vorprog/go-api/server"
	"github.com/vorprog/go-api/util"
)

func main() {
	util.Log("Loaded main module.")
	if os.Getenv("SOPS_FILE_URL") != "" {
		var sopsError = util.SetEnvironmentFromSopsURL()
		if sopsError != nil {
			util.Log(sopsError)
			return
		}
	}

	awsIdentity, err := util.GetAwsIdentity()

	if err != nil {
		log.Println(err)
	} else {
		util.Log(awsIdentity)
	}

	go util.Monitor()

	port, _ := lo.Coalesce(os.Getenv("APP_SERVER_PORT"), "8080")
	err = server.Start(port)

	if err != nil {
		log.Fatal(err)
	}
}
