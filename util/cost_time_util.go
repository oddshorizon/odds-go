package util

import (
	"fmt"
	"time"
)

//
//  CostTime
//  @Description: 统计耗时
//  @param start
//  @return int64
//
func CostTime(start time.Time) int64 {
	return time.Since(start).Milliseconds()
}

// 耗时统计
func PrintCostTime() func() {
	start := time.Now()
	return func() {
		ct := time.Since(start).Milliseconds()
		fmt.Sprintf("cost time %dms", ct)
	}
}
