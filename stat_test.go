package main

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/google/gofuzz"
)

func TestDecodeFile(t *testing.T) {
	fuzzy := fuzz.New()

	checkBadFile := func() {
		badFileReader := func(_ readFile) ([]byte, error) {
			var junk string
			fuzzy.Fuzz(&junk)
			return []byte(junk), nil
		}

		var filePath string
		fuzzy.Fuzz(&filePath)
		_, println := decodeFile(badFileReader, filePath, emptyStat)
		if !strings.Contains(string(println), filePath) || !strings.Contains(string(println), "Parse failure") {
			t.Fatal("Bad json should have thrown an error", println, filePath)
		}
	}

	// validReader := func(_ readFile) ([]byte, error) {
	// 	kind := getRandomFrom([]string{alarmKind, doorKind, imgKind})
	// 	return []byte("{\"Type\":\"" + kind + "\"}"), nil
	// }

	forN(100, checkBadFile)
}

func TestUpdateStat(t *testing.T) {
	rand.Seed(time.Now().Unix())
	proof := func() {
		kind := getRandomFrom(
			[]string{alarmKind, doorKind, imgKind, "crap"})
		updatedStat := updateStat(kind, emptyStat)
		if updatedStat == emptyStat && kind != "crap" {
			t.Fatal("stat failed to update with valid option")
		}
		if updatedStat != emptyStat && kind == "crap" {
			t.Fatal("stat updated when invalid option was passed")
		}
		if updatedStat != emptyStat && kind != "crap" && (updatedStat.alarmCnt+updatedStat.doorCnt+updatedStat.imgCnt) != 1 {
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
