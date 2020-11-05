package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Path is: "+os.Getenv("PATH"))
}
