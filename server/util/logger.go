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
		var messageString string
		jsonMessage, jsonMarshalError := json.Marshal(message)
		if jsonMarshalError != nil {
			messageString = fmt.Sprint(jsonMarshalError)
		} else {
			messageString = string(jsonMessage)
		}
		callerPrefix := file + "(" + strconv.Itoa(line) + ") "
		log.Println(callerPrefix + messageString)
	}
}
