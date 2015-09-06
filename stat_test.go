package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/gofuzz"
)

func choose(min, max int) int {
	return rand.Intn(max-min) + min
}

func getRandomFrom(xs []string) string {
	return xs[choose(0, len(xs))]
}

func TestUpdateStat(t *testing.T) {
	options := []string{alarmKind, doorKind, imgKind, "crap"}
	rand.Seed(time.Now().Unix())
	proof := func() {
		option := getRandomFrom(options)
		updatedStat := updateStat(option, emptyStat)
		if updatedStat == emptyStat && option != "crap" {
			t.Fatal("stat failed to update with valid option")
		}
		if updatedStat != emptyStat && option == "crap" {
			t.Fatal("stat updated when invalid option was passed")
		}
		if updatedStat != emptyStat && option != "crap" && (updatedStat.alarmCnt+updatedStat.doorCnt+updatedStat.imgCnt) != 1 {
			t.Fatal("stat updated incorrectly")
		}
	}
	forN(100, proof)
}

func TestCalcAvg(t *testing.T) {
	fuzzy := fuzz.New()
	proof := func() {
		var doorCnt int
		var alarmCnt int
		var imgCnt int
		var duration int
		fuzzy.Fuzz(&doorCnt)
		fuzzy.Fuzz(&alarmCnt)
		fuzzy.Fuzz(&imgCnt)
		fuzzy.Fuzz(&duration)

		sampleAvg := func() int64 {
			total := doorCnt + imgCnt + alarmCnt
			if total == 0 {
				return 0
			}
			return int64(duration / total)
		}

		avgProcessingTime :=
			calcAvg(state{
				stat{doorCnt, imgCnt, alarmCnt, 0},
				time.Duration(duration)}).avgProcessingTime.Nanoseconds()

		if avgProcessingTime != sampleAvg() {
			t.Fatal("average computation is incorrect with")
		}

	}
	forN(100, proof)
}
