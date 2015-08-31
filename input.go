package main

import (
	"encoding/json"
	"time"
)

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

func decodeAlarm(j []byte) (alarm, error) {
	var a alarm
	err := json.Unmarshal(j, &a)
	return a, err
}

func decodeDoor(j []byte) (door, error) {
	var a door
	err := json.Unmarshal(j, &a)
	return a, err
}
func decodeImg(j []byte) (img, error) {
	var a img
	err := json.Unmarshal(j, &a)
	return a, err
}
