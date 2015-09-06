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
	forN(100, func() {
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
	})
}

func TestCalcAvg(t *testing.T) {
	var doorCnt int
	var alarmCnt int
	var imgCnt int
	var durration int
	var sampleAvg int64
	fuzzy := fuzz.New()
	forN(100, func() {
		fuzzy.Fuzz(&doorCnt)
		fuzzy.Fuzz(&alarmCnt)
		fuzzy.Fuzz(&imgCnt)
		fuzzy.Fuzz(&durration)

		s := state{
			stat{doorCnt, imgCnt, alarmCnt, 0},
			time.Duration(durration)}

		total := doorCnt + imgCnt + alarmCnt

		if total == 0 {
			sampleAvg = 0
		} else {
			sampleAvg = int64(durration / total)
		}

		avgProcessingTime :=
			calcAvg(s).avgProcessingTime.Nanoseconds()

		if avgProcessingTime != sampleAvg {
			t.Fatal("average computation is incorrect with")
		}
	})
}
