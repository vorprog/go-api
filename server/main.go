package main

import (
	"flag"

	"github.com/richardsnider/golang/server/server"
	"github.com/richardsnider/golang/server/util"
)

var defaultPort = "8080"
var port = *flag.String("port", defaultPort, "port to listen on")

func main() {
	util.Log("Loaded main module.")
	awsIdentity := util.GetAwsIdentity()
	util.Log(awsIdentity)

	go util.Monitor()
	server.Start(port)
}
