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
	proof := func() {
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
	forN(100, proof)
}

// func TestCollect(t *testing.T) {
// 	reader := func(readFile) ([]byte, error){
//
// 	}
// }
