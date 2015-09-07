package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/go-fsnotify/fsnotify"
	"github.com/google/gofuzz"
)

func TestWhenCreation(t *testing.T) {
	var event fsnotify.Event
	fuzzy := fuzz.New()
	in := make(chan fsnotify.Event)
	out := whenCreation(in)
	write := fsnotify.Write
	create := fsnotify.Create

	prop := func() {
		fuzzy.Fuzz(&event)
		in <- event

		go func() {
			for output := range out {
				if !(output.Op&create == create || output.Op&write == write) {

					t.Fatal(
						"Something other than create and write made it through",
						output)
				}
			}
		}()
	}

	check(prop)
}

func TestCollect(t *testing.T) {
	rand.Seed(time.Now().Unix())
	fuzzy := fuzz.New()
	events := make(chan fsnotify.Event)
	ticks := make(chan time.Time)
	done := make(chan bool)
	count := 0

	dumbyReader := func(readFile) ([]byte, error) {
		return []byte{}, nil
	}

	decoder := func(
		_ fileReader,
		_ string,
		stat stat) (stat, println) {

		switch choose(0, 3) {

		case 0:
			stat.alarmCnt++

		case 1:
			stat.doorCnt++

		case 2:
			stat.imgCnt++
		}

		return stat, ""
	}

	var filePath string
	fuzzy.Fuzz(&filePath)

	out := collect(
		dumbyReader, decoder, events, ticks)

	go func() {
		finalstat := <-out
		if finalstat.alarmCnt+finalstat.doorCnt+finalstat.imgCnt != count {
			t.Fatal("stat did not increment per tick")
		}
		go func() {
			done <- true
		}()
	}()

	prop := func() {
		var event fsnotify.Event
		var tick time.Time

		forN(choose(0, 1000), func() {
			fuzzy.Fuzz(&event)
			// count++
			events <- event
		})

		fuzzy.Fuzz(&tick)
		ticks <- tick
		<-done
	}

	prop()
	// prop()

}
