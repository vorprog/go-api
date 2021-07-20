package util

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func Log(message ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		callerPrefix := file + "(" + strconv.Itoa(line) + ") "
		messageString := fmt.Sprint(message...)
		os.Stdout.WriteString(callerPrefix + messageString)
	}
}

func LogError(message ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		callerPrefix := file + "(" + strconv.Itoa(line) + ") "
		messageString := fmt.Sprint(message...)
		os.Stderr.WriteString(callerPrefix + messageString)
	}
}
