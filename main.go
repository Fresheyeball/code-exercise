package main

import (
	"log"

	fsnotify "github.com/go-fsnotify/fsnotify"
)

type stat struct {
	doorCnt           int
	ImgCnt            int
	alarmCnt          int
	avgProcessingTime int
}

func fatality(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func watchInput() *fsnotify.Watcher {
	watcher, watcherr := fsnotify.NewWatcher()
	fatality(watcherr)

	adderr := watcher.Add("input/")
	fatality(adderr)

	return watcher
}

func retrieve(watcher *fsnotify.Watcher) <-chan fsnotify.Event {
	events := make(chan fsnotify.Event)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				events <- event
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()
	return events
}

func collect(<-chan fsnotify.Event) <-chan stat {
	return make(chan stat)
}

func main() {

	watcher := watchInput()
	log.Println(<-retrieve(watcher))

	// cleanup before you die
	defer watcher.Close()
	<-make(chan bool)
}
