package main

import (
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

func whenCreation(events <-chan fsnotify.Event) <-chan fsnotify.Event {
	out := make(chan fsnotify.Event)
	write := fsnotify.Write
	create := fsnotify.Create
	go func() {
		for event := range events {
			switch {
			case event.Op&create == create:
				out <- event
			case event.Op&write == write:
				out <- event
			}
		}
	}()
	return out
}

func collectOn(events <-chan (fsnotify.Event), ticker <-chan (time.Time)) <-chan (stat) {
	stats := make(chan stat)
	state := emptyState
	go func() {
		for {
			select {
			case event := <-events:
				eventTime := time.Now()
				state.stat = decodeFile(event.Name, state.stat)
				state.duration = state.duration + time.Since(eventTime)
			case <-ticker:
				go func() {
					stats <- calcAvg(state)
					state = emptyState
				}()
			}
		}
	}()
	return stats
}

func main() {
	w := logErrors(watchInput("input/"))
	printStats(collectOn(whenCreation(w.watcher.Events), time.NewTicker(time.Second).C))
	death.NewDeath(SYS.SIGINT, SYS.SIGTERM).WaitForDeath(w)
}
