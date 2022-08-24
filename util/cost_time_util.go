package util

import (
	"time"
)

// 耗时统计
func CostTime() func() int64 {
	start := time.Now()
	return func() int64 {
		return time.Since(start).Milliseconds()
	}
}
