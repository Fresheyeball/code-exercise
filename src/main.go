package main

import (
	"log"
	SYS "syscall"

	death "github.com/vrecan/death"
)

func main() {
	w := watchInput("/input")

	log.Println("event:", <-events(w))

	death.NewDeath(SYS.SIGINT, SYS.SIGTERM).WaitForDeath(w)
}
