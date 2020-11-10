package main

import (
	"fmt"

	"github.com/richardsnider/golang/hello_world/server"
)

func main() {
	fmt.Println("Started main package.")
	server.Start()
}
