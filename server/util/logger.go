package util

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
)

func Log(message ...interface{}) {
	log.SetFlags(0)
	_, file, line, ok := runtime.Caller(1)
	if ok {
		callerPrefix := file + "(" + strconv.Itoa(line) + ") "
		messageString := fmt.Sprint(message...)
		log.Println(callerPrefix + messageString)
	}
}
