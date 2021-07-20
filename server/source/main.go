package main

import (
	"flag"

	"github.com/richardsnider/golang/server/source/server"
	"github.com/richardsnider/golang/server/source/util"
)

var defaultPort = "8080"
var port = flag.String("port", defaultPort, "port to listen on")

func main() {
	util.Log("Started main package.")
	server.Start(port)
}
