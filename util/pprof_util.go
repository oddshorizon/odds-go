package util

import (
	"fmt"
	"os"
	"runtime/pprof"
)

var cpuProfileFile os.File
var memProfileFile os.File

func LaunchPprof(pprofEnabled bool, dir string) error {
	if !pprofEnabled {
		return nil
	}
	// start cpu profile
	cpuProfileFilePath := fmt.Sprintf("%s/cpu.profile", dir)
	os.Remove(cpuProfileFilePath)
	cpuProfileFile, err := os.Create(cpuProfileFilePath)
	if err != nil {
		return err
	}
	if err := pprof.StartCPUProfile(cpuProfileFile); err != nil {
		return err
	}
	// start mem profile
	memProfileFilePath := fmt.Sprintf("%s/mem.profile", dir)
	os.Remove(memProfileFilePath)
	memProfileFile, err := os.Create(memProfileFilePath)
	if err != nil {
		return err
	}
	if err := pprof.WriteHeapProfile(memProfileFile); err != nil {
		return err
	}
	return nil
}

func StopPprof() {
	if nil != &cpuProfileFile {
		pprof.StopCPUProfile()
		cpuProfileFile.Close()
	}
	if nil != &memProfileFile {
		memProfileFile.Close()
	}
}
