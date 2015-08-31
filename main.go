package main

import (
	"log"
	SYS "syscall"
	"time"

	"github.com/go-fsnotify/fsnotify"
	death "github.com/vrecan/death"
)

type state struct {
	stat     stat
	duration time.Duration
}

var emptyState = state{emptyStat, 0}

func collectOn(events <-chan (fsnotify.Event), ticker <-chan (time.Time)) <-chan (stat) {
	stats := make(chan stat)
	state := emptyState

	go func() {
		for event := range events {
			log.Println("event", event)
			eventTime := time.Now()
			state.stat = process(event.Name, state.stat)
			state.duration = state.duration + time.Since(eventTime)
		}
	}()

	go func() {
		for range ticker {
			log.Println("tick")
			go func() {
				newState := calcAvg(state)
				log.Println("send", newState)
				stats <- newState
				state = emptyState
			}()
		}
	}()

	return stats
}

func main() {
	w := logErrors(watchInput("input/"))

	printStats(collectOn(w.watcher.Events, time.NewTicker(time.Second).C))

	death.NewDeath(SYS.SIGINT, SYS.SIGTERM).WaitForDeath(w)
}
