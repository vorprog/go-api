package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

var appVersion = `1.0`
var buildDateVersionLinkerFlag string
var buildCommitLinkerFlag string

type buildMetaData struct {
	Version          string `json:"Version"`
	BuildGitCommit   string `json:"BuildGitCommit"`
	BuildDateVersion string `json:"BuildDateVersion"`
}

func jsEchoFunction() js.Func {
	jsFunction := func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no arguments passed"
		}
		inputString := args[0].String()
		fmt.Println(inputString)
		return inputString
	}

	return js.FuncOf(jsFunction)
}

func main() {
	wasmBuildMetaData := buildMetaData{
		Version:          appVersion,
		BuildGitCommit:   buildCommitLinkerFlag,
		BuildDateVersion: buildDateVersionLinkerFlag,
	}

	metaDataBytes, _ := json.Marshal(wasmBuildMetaData)
	metaData := string(metaDataBytes)
	fmt.Println(metaData)

	htmlDom := js.Global().Get("document")
	p := htmlDom.Call("createElement", "p")
	p.Set("innerHTML", metaData)
	htmlDom.Get("body").Call("appendChild", p)

	js.Global().Set("echo", jsEchoFunction())
}
