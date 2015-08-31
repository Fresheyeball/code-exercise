package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type stat struct {
	doorCnt           int
	imgCnt            int
	alarmCnt          int
	avgProcessingTime time.Duration
}

func encodeStat(s stat) []byte {
	j, _ := json.Marshal(s)
	return j
}

var emptyStat = stat{0, 0, 0, 0}

func process(filePath string, stat stat) stat {
	contents, _ := ioutil.ReadFile(filePath)

	switch {
	case isError(decodeAlarm(contents)):
		stat.alarmCnt++
	case isError(decodeDoor(contents)):
		stat.doorCnt++
	case isError(decodeImg(contents)):
		stat.imgCnt++
	}

	return stat
}

func calcAvg(state state) stat {
	stat := state.stat
	total := int64(stat.doorCnt + stat.alarmCnt + stat.doorCnt)
	stat.avgProcessingTime = time.Duration(safeDiv(state.duration.Nanoseconds(), total))
	return stat
}

func printStat(stat stat) {
	avgProcessingTime := stat.avgProcessingTime.Nanoseconds() / time.Millisecond.Nanoseconds()
	fmt.Printf("DoorCnt: %d, ImgCnt: %d, AlarmCnt: %d, avgProcessingTime: %dms \n", stat.doorCnt, stat.imgCnt, stat.alarmCnt, avgProcessingTime)
}

func printStats(stats <-chan (stat)) {
	go func() {
		for {
			printStat(<-stats)
		}
	}()
}
