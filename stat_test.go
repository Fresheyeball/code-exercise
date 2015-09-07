package main

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/google/gofuzz"
)

func fuzzyStat() stat {
	getRand := func() int {
		return choose(0, 1000000000000000000)
	}

	return stat{
		getRand(),
		getRand(),
		getRand(),
		time.Duration(getRand())}
}

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

		_, println := decodeFile(badFileReader, filePath, fuzzyStat())

		if !strings.Contains(string(println), filePath) ||
			!strings.Contains(string(println), "Parse failure") {

			t.Fatal("Bad json should have logged an error", println, filePath)
		}
	}

	checkBadKind := func() {
		var kind string
		fuzzy.Fuzz(&kind)

		// this test is not about invalid json
		// or bad decoding, but having the wrong kind
		if strings.ContainsAny(kind, `\"`) || kind == "" {
			return
		}

		badKindReader := func(_ readFile) ([]byte, error) {
			return []byte("{\"Type\":\"" + kind + "\"}"), nil
		}

		var filePath string
		fuzzy.Fuzz(&filePath)

		_, println := decodeFile(badKindReader, filePath, fuzzyStat())

		if !strings.Contains(string(println), filePath) ||
			!strings.Contains(string(println), kind) ||
			!strings.Contains(string(println), "Parse successful") {

			runPrintln(println)

			t.Fatal("Unknown kind should have logged an error", kind)
		}
	}

	checkValid := func() {
		kind := getRandomFrom([]string{alarmKind, doorKind, imgKind})
		validReader := func(_ readFile) ([]byte, error) {
			return []byte("{\"Type\":\"" + kind + "\"}"), nil
		}

		var filePath string
		fuzzy.Fuzz(&filePath)

		_, println := decodeFile(validReader, filePath, fuzzyStat())

		if string(println) != "" {
			t.Fatal("Valid json logged an error")
		}
	}

	forN(checkSize, checkBadFile)
	forN(checkSize, checkBadKind)
	forN(checkSize, checkValid)
}

func TestUpdateStat(t *testing.T) {
	rand.Seed(time.Now().Unix())

	check := func() {
		kind := getRandomFrom(
			[]string{alarmKind, doorKind, imgKind, "crap"})
		updatedStat := updateStat(kind, emptyStat)

		if updatedStat == emptyStat &&
			kind != "crap" {

			t.Fatal("stat failed to update with valid option")
		}

		if updatedStat != emptyStat &&
			kind == "crap" {

			t.Fatal("stat updated when invalid option was passed")
		}

		if updatedStat != emptyStat &&
			kind != "crap" &&
			(updatedStat.alarmCnt+
				updatedStat.doorCnt+
				updatedStat.imgCnt) != 1 {

			t.Fatal("stat updated incorrectly")
		}
	}

	forN(checkSize, check)
}

func TestCalcAvg(t *testing.T) {
	fuzzy := fuzz.New()

	check := func() {
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

	forN(checkSize, check)
}
