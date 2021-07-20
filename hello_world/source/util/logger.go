package util

import (
	"log"
	"runtime"
	"strconv"
)

func Log(message string) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		callerPrefix := file + "(" + strconv.Itoa(line) + ") "
		log.Println(callerPrefix + message)
	}
}
