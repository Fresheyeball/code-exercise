package main

import (
	"log"
	SYS "syscall"
	"time"

	"github.com/go-fsnotify/fsnotify"
	death "github.com/vrecan/death"
)

func collect(events <-chan (fsnotify.Event)) <-chan (stat) {
	stats := make(chan stat)
	eventTime := time.Now()
	event := <-events
	log.Println("--->", eventTime, event.Name)

	return stats
}

func main() {
	w := logErrors(watchInput("input/"))

	log.Println("event:", <-w.watcher.Events)

	death.NewDeath(SYS.SIGINT, SYS.SIGTERM).WaitForDeath(w)
}
