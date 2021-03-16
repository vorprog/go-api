package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var directoryPath = flag.String("directory", "/sys", "directory to list")

func main() {
	fmt.Fprintln(os.Stdout, "Path is: "+os.Getenv("PATH"))
	fmt.Fprintln(os.Stdout, "Current directory is: "+os.Getenv("PWD"))

	err := filepath.Walk(*directoryPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
