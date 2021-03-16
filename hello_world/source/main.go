package main

import (
	"flag"
	"fmt"

	"github.com/richardsnider/golang/hello_world/server"
)

var defaultPort = "8080"
var port = flag.String("port", defaultPort, "port to listen on")

func main() {
	fmt.Println("Started main package.")
	server.Start(port)
}
