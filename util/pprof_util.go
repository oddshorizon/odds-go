package util

import (
	"fmt"
	"github.com/juqiukai/glog"
	"os"
	"runtime/pprof"
	"time"
)

var mainPprofEnabled bool
var cpuProfileFile os.File
var memProfileFile os.File

func LaunchPprof(pprofEnabled bool, dir string) error {
	mainPprofEnabled = pprofEnabled
	if !pprofEnabled {
		return nil
	}
	glog.Infof("begin start pprof ... ")
	// start cpu profile
	cpuProfileFilePath := fmt.Sprintf("%s/cpu.pprof", dir)
	os.Remove(cpuProfileFilePath)
	cpuProfileFile, err := os.Create(cpuProfileFilePath)
	if err != nil {
		return err
	}
	if err := pprof.StartCPUProfile(cpuProfileFile); err != nil {
		return err
	}
	// start mem profile
	memProfileFilePath := fmt.Sprintf("%s/mem.pprof", dir)
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

func StopPprof(waitSeconds int) {
	if !mainPprofEnabled {
		return
	}
	go func() {
		time.Sleep(time.Duration(waitSeconds) * time.Second)
		if nil != &cpuProfileFile {
			pprof.StopCPUProfile()
			cpuProfileFile.Close()
		}
		if nil != &memProfileFile {
			memProfileFile.Close()
		}
		glog.Infof("stop pprof success - waitSeconds=%d", waitSeconds)
	}()
}
