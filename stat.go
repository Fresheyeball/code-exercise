package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type stat struct {
	doorCnt           int
	imgCnt            int
	alarmCnt          int
	avgProcessingTime time.Duration
}

var emptyStat = stat{0, 0, 0, 0}

func process(filePath string, stat stat) stat {
	decoded, err := decode(attemptGet(ioutil.ReadFile(filePath)).([]byte))
	if err != nil {
		log.Println("Parse error with: " + filePath)
		return stat
	}
	switch decoded.Kind {
	case alarmKind:
		stat.alarmCnt++
	case doorKind:
		stat.doorCnt++
	case imgKind:
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
