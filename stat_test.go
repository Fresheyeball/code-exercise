package main

import (
	"math/rand"
	"strconv"
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

	propBadFile := func() {
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

	propBadKind := func() {
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

	propValid := func() {
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

	check(propBadFile)
	check(propBadKind)
	check(propValid)
}

func TestUpdateStat(t *testing.T) {
	rand.Seed(time.Now().Unix())

	prop := func() {
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

	check(prop)
}

func TestCalcAvg(t *testing.T) {
	fuzzy := fuzz.New()

	prop := func() {
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

	check(prop)
}

func TestRenderStat(t *testing.T) {
	prop := func() {
		stat := fuzzyStat()
		rendered := renderStat(stat)
		milliseconds :=
			stat.avgProcessingTime.Nanoseconds() /
				time.Millisecond.Nanoseconds()
		contained := func(name string, i int) {
			if !(strings.Contains(rendered, strconv.Itoa(i)) &&
				strings.Contains(rendered, name)) {
				t.Fatal(name + " did not appear in rendered stat")
			}
		}

		contained("AlarmCnt", stat.alarmCnt)
		contained("DoorCnt", stat.doorCnt)
		contained("ImgCnt", stat.imgCnt)
		contained("AvgProcessingTime", int(milliseconds))
	}

	check(prop)
}
