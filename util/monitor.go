package util

import (
	"bytes"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func Monitor() {
	var cpuWriter bytes.Buffer
	runtime.SetCPUProfileRate(400)
	err := pprof.StartCPUProfile(&cpuWriter)

	if err != nil {
		Log(err)
		os.Exit(1)
	}

	for {
		time.Sleep(2 * time.Second)

		memory := &runtime.MemStats{}
		runtime.ReadMemStats(memory)

		Log(map[string]interface{}{
			"cpu":    cpuWriter.String(),
			"memory": memory.Alloc,
		})
	}
}
