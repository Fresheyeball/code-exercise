package main

import (
	"encoding/json"
	"time"
)

type alarm struct {
	Date  time.Time `json:"Date"`
	Name  string    `json:"name"`
	Floor int       `json:"floor"`
	Room  int       `json:"room"`
}

type door struct {
	Date time.Time `json:"Date"`
	Open bool      `json:"open"`
}

type img struct {
	Date  time.Time `json:"Date"`
	Size  int       `json:"size"`
	Bytes []byte    `json:"bytes"`
}

func decodeAlarm(j []byte) (alarm, error) {
	a := alarm{}
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
