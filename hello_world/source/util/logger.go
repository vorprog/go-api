package util

import (
	"encoding/json"
	"log"
	"runtime/debug"
)

func Log(message string) {
	logJSON, err := json.Marshal(map[string]interface{}{
		"stack":   string(debug.Stack()),
		"message": "port is " + message,
	})
	if err != nil {
		log.Fatalf("Unable to encode JSON in logger")
	}

	log.Println(string(logJSON))
}
