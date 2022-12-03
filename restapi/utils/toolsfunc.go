package utils

import (
	"fmt"
	"time"
)

// GetNowFormatTodayTime 获取当天时间
func GetNowFormatTodayTime() string {
	now := time.Now()
	dateStr := fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()), now.Day())
	return dateStr
}
