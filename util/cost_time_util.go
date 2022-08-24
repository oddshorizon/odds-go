package util

import (
	"time"
)

//@brief：耗时统计函数
func CostTime() func() int64 {
	start := time.Now()
	return func() int64 {
		return time.Since(start).Milliseconds()
	}
}
