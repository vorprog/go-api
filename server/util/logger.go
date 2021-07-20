package util

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
)

func Log(message ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		callerPrefix := file + "(" + strconv.Itoa(line) + ") "
		messageString := fmt.Sprint(message...)
		log.Println(callerPrefix + messageString)
	}
}
