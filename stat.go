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

func decodeFile(filePath string, stat stat) stat {
	decoded, err := decode(attemptGet(
		ioutil.ReadFile(filePath)).([]byte))

	if err != nil {
		log.Println("Parse failure in file: "+filePath+" With: ", err)
		return stat
	}

	updatedStat := updateStat(decoded.Kind, stat)

	if updatedStat == stat {
		log.Println(
			"Parse successful but not a known type in file: "+filePath+" Found: ",
			decoded.Kind)
	}

	return updatedStat
}

func updateStat(kind string, stat stat) stat {
	switch kind {
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
	total := int64(stat.doorCnt + stat.alarmCnt + stat.imgCnt)
	stat.avgProcessingTime =
		time.Duration(safeDiv(state.duration.Nanoseconds(), total))
	return stat
}

func renderStat(stat stat) string {
	avgProcessingTime :=
		stat.avgProcessingTime.Nanoseconds() / time.Millisecond.Nanoseconds()
	return fmt.Sprintf(
		"DoorCnt: %d, ImgCnt: %d, AlarmCnt: %d, avgProcessingTime: %dms",
		stat.doorCnt,
		stat.imgCnt,
		stat.alarmCnt,
		avgProcessingTime)
}

func printStats(stats <-chan (stat)) {
	go func() {
		for {
			log.Println(renderStat(<-stats))
		}
	}()
}
