package util

import (
	"fmt"
	"github.com/juqiukai/glog"
	"os"
	"runtime/pprof"
	"time"
)

var mainPprofEnabled bool
var cpuProfileFile *os.File
var memProfileFile *os.File

func LaunchPprof(pprofEnabled bool, dir string) error {
	mainPprofEnabled = pprofEnabled
	if !pprofEnabled {
		return nil
	}
	glog.Infof("begin start pprof ... ")
	// start cpu profile
	cpuProfileFilePath := fmt.Sprintf("%s/cpu.pprof", dir)
	os.Remove(cpuProfileFilePath)
	f1, err := os.Create(cpuProfileFilePath)
	if err != nil {
		return err
	}
	cpuProfileFile = f1
	if err := pprof.StartCPUProfile(cpuProfileFile); err != nil {
		return err
	}
	// start mem profile
	memProfileFilePath := fmt.Sprintf("%s/mem.pprof", dir)
	os.Remove(memProfileFilePath)
	f2, err := os.Create(memProfileFilePath)
	if err != nil {
		return err
	}
	memProfileFile = f2
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
		if nil != memProfileFile {
			pprof.WriteHeapProfile(memProfileFile)
			memProfileFile.Close()
		}
		glog.Infof("stop pprof success - waitSeconds=%d", waitSeconds)
	}()
}
