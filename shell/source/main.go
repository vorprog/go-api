package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// TODO: process ~/.profile bash code
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(os.Getenv("PS1"))

		// TODO: account for multi-line input using backslash character
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		input = strings.TrimSuffix(input, "\n")

		if input == "exit" {
			os.Exit(0)
		}

		err = execInput(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	arguments := strings.Split(input, " ")
	command := exec.Command(arguments[0], arguments[1:]...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	return command.Run()
}
