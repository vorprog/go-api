package util

import (
	"runtime"
	"time"
)

func Monitor() {
	for {
		time.Sleep(2 * time.Second)

		memory := &runtime.MemStats{}
		runtime.ReadMemStats(memory)

		Log(map[string]interface{}{
			"cpu":       runtime.NumCPU(),
			"memory":    memory.Alloc,
			"goRoutine": runtime.NumGoroutine(),
		})
	}
}
