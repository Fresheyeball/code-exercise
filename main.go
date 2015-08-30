package main

import (
	"io/ioutil"
	"log"
	SYS "syscall"
	"time"

	"github.com/go-fsnotify/fsnotify"
	death "github.com/vrecan/death"
)

func process(filePath string, initial stat) stat {
	contents, _ := ioutil.ReadFile(filePath)
	a, _ := decodeAlarm(contents)
	d, _ := decodeDoor(contents)
	i, _ := decodeImg(contents)
	switch {
	case a != nil:
		initial.alarmCnt++
	case d != nil:
		initial.doorCnt++
	case i != nil:
		initial.imgCnt++
	}

	return emptyStat
}

func collectOn(events <-chan (fsnotify.Event), ticker <-chan (time.Time)) <-chan (stat) {
	stats := make(chan stat)
	state := emptyStat

	go func() {
		for {
			select {
			case event := <-events:
				eventTime := time.Now()

				log.Println("--->", eventTime, event.Name)
			case <-ticker:
				stats <- state
				state = emptyStat
			}
		}
	}()

	return stats
}

func main() {
	w := logErrors(watchInput("input/"))

	go collectOn(w.watcher.Events, time.NewTicker(time.Second).C)

	death.NewDeath(SYS.SIGINT, SYS.SIGTERM).WaitForDeath(w)
}
