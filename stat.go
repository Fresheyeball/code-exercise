package main

import (
	"fmt"
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

func decodeFile(
	fileReader fileReader,
	filePath string,
	stat stat) (stat, println) {

	decoded, err := decode(attemptGet(
		fileReader(readFile(filePath))).([]byte))

	if err != nil {
		return stat, println(
			"Parse failure in file: " +
				string(filePath) +
				" With: " +
				err.Error())
	}

	updatedStat := updateStat(decoded.Kind, stat)
	log.Println("---> ", updatedStat)
	handleBadKind := func() println {
		if updatedStat == stat {
			return println(
				"Parse successful but not a known type in file: " +
					string(filePath) +
					" Found: " +
					decoded.Kind)
		}
		return println("")
	}

	return updatedStat, handleBadKind()
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
		"DoorCnt: %d, "+
			"ImgCnt: %d, "+
			"AlarmCnt: %d, "+
			"avgProcessingTime: %dms",
		stat.doorCnt,
		stat.imgCnt,
		stat.alarmCnt,
		avgProcessingTime)
}

func printStats(stats <-chan stat) {
	for {
		log.Println(renderStat(<-stats))
	}
}
