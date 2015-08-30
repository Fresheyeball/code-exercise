package main

import (
	"log"

	"github.com/go-fsnotify/fsnotify"
)

type watcher struct {
	watcher *fsnotify.Watcher
}

// newWatcher : IO Watcher
func newWatcher() watcher {
	w, err := fsnotify.NewWatcher()
	fatality(err)
	return watcher{w}
}

// watchInput : String -> IO Watcher
func watchInput(input string) watcher {
	w := newWatcher()
	fatality(w.watcher.Add(input))
	return w
}

func (w watcher) Close() {
	fatality(w.watcher.Close())
}

func events(w watcher) <-chan fsnotify.Event {
	x := make(chan fsnotify.Event)
	go func() {
		for {
			select {
			case event := <-w.watcher.Events:
				x <- event
			case err := <-w.watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()
	return x
}
