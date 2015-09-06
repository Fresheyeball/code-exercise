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

func collect(
	decoder func(string, stat) stat,
	events <-chan fsnotify.Event,
	ticker <-chan time.Time) <-chan stat {

	stats := make(chan stat)
	go func() {
		state := emptyState
		for {
			select {
			case event := <-events:
				eventTime := time.Now()
				state.stat = decoder(event.Name, state.stat)
				state.duration += time.Since(eventTime)
			case <-ticker:
				stats <- calcAvg(state)
				state = emptyState
			}
		}
	}()
	return stats
}

func main() {
	w := logErrors(watchInput("input/"))
	printStats(collect(
		decodeFile,
		whenCreation(w.watcher.Events),
		time.NewTicker(time.Second).C))
	death.NewDeath(SYS.SIGINT, SYS.SIGTERM).WaitForDeath(w)
}
