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

// func TestCollectOn(t *testing.T) {
// 	var event fsnotify.Event
// 	var tick time.Time
// 	fuzzy := fuzz.New()
// 	events := make(chan fsnotify.Event)
// 	ticks := make(chan time.Time)
// 	out := collectOn(events, ticks)
// 	forN(100, func() {
// 		fuzzy.Fuzz(&event)
// 		fuzzy.Fuzz(&tick)
// 		events <- event
// 		ticks <- tick
// 		go func() {
// 			for {
// 				select {}
// 			}
// 		}()
// 	})
// }
