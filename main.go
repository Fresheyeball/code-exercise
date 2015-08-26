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

func collect(watcher *fsnotify.Watcher) <-chan fsnotify.Event {
	x := make(chan fsnotify.Event)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				x <- event
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()
	return x
}

func main() {

	watcher := watchInput()

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	// cleanup before you die
	defer watcher.Close()
	<-make(chan bool)
}
