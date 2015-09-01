package main

import (
	"encoding/json"
	"errors"
	"time"
)

const (
	alarmKind string = "alarm"
	doorKind  string = "door"
	imgKind   string = "img"
)

type alarm struct {
	Kind  string    `json:"Type"`
	Date  time.Time `json:"Date"`
	Name  string    `json:"name"`
	Floor int       `json:"floor"`
	Room  int       `json:"room"`
}

type door struct {
	Kind string    `json:"Type"`
	Date time.Time `json:"Date"`
	Open bool      `json:"open"`
}

type img struct {
	Kind  string    `json:"Type"`
	Date  time.Time `json:"Date"`
	Size  int       `json:"size"`
	Bytes []byte    `json:"bytes"`
}

func wrongKind(kind string) error {
	return errors.New("parsed to " + kind + ", but the type in json is wrong")
}

func decodeAlarm(j []byte) (alarm, error) {
	var a alarm
	err := json.Unmarshal(j, &a)
	if a.Kind != alarmKind {
		return a, wrongKind(alarmKind)
	}
	return a, err
}

func decodeDoor(j []byte) (door, error) {
	var a door
	err := json.Unmarshal(j, &a)
	if a.Kind != doorKind {
		return a, wrongKind(doorKind)
	}
	return a, err
}

func decodeImg(j []byte) (img, error) {
	var a img
	err := json.Unmarshal(j, &a)
	if a.Kind != imgKind {
		return a, wrongKind(imgKind)
	}
	return a, err
}
