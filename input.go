package main

import "time"

type alarm struct {
	date  time.Time
	name  string
	floor int
	room  int
}

type door struct {
	date time.Time
	open bool
}

type img struct {
	date  time.Time
	size  int
	bytes []byte
}
