package main

import (
	"testing"

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
	forN(100, func() {
		fuzzy.Fuzz(&event)
		in <- event
		go func() {
			for output := range out {
				if !(output.Op&create == create || output.Op&write == write) {
					t.Fatal("Something other than create and write made it through", output)
				}
			}
		}()
	})
}

func fuzzyDecoder(_ string, stat stat) stat {
	return updateStat(getRandomFrom([]string{alarmKind, doorKind, imgKind}), stat)
}

// func TestCollectOn(t *testing.T) {
// 	var event fsnotify.Event
// 	var tick time.Time
// 	fuzzy := fuzz.New()
// 	events := make(chan fsnotify.Event)
// 	ticks := make(chan time.Time)
// 	out := collect(fuzzyDecoder, events, ticks)
//
// 	forN(100, func() {
// 		count := 0
//
// 		forN(choose(0, 100), func() {
// 			fuzzy.Fuzz(&event)
// 			count++
// 			events <- event
// 		})
// 		go func() {
// 			fuzzy.Fuzz(&tick)
// 			ticks <- tick
// 		}()
// 		go func() {
// 			for output := range out {
// 				log.Println("count", count, "\toutput", output)
// 				if count != output.doorCnt+output.imgCnt+output.alarmCnt {
// 					t.Fatal("counts dont add up")
// 				}
// 				count = 0
// 			}
// 		}()
// 	})
//
// }
