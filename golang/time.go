package main

import (
	"fmt"
	"time"
)

// 计算时间差

func TimeDiff(startTime, stopTime string) int64 {
	var hour int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", stopTime, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 3600
		return hour
	} else {
		return hour
	}
}

func main() {
	fmt.Println(TimeDiff("2016-09-10 13:00:00", "2016-09-11 15:00:00"))
}
