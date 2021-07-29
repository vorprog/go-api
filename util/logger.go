package util

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"strconv"
)

func Log(message ...interface{}) {
	log.SetFlags(0)
	_, file, line, ok := runtime.Caller(1)
	if ok {
		jsonMessage, jsonMarshalError := json.Marshal(map[string]interface{}{
			"source_code": file + "(" + strconv.Itoa(line) + ") ",
			"message":     message,
		})

		if jsonMarshalError != nil {
			log.Println(fmt.Sprint(jsonMarshalError))
		} else {
			log.Println(string(jsonMessage))
		}
	}
}
