package main

import (
	"encoding/json"
	"errors"
)

const (
	alarmKind string = "alarm"
	doorKind  string = "door"
	imgKind   string = "img"
)

type input struct {
	Kind string `json:"Type"`
}

func decode(j []byte) (input, error) {
	var i input
	err := json.Unmarshal(j, &i)

	if i.Kind == "" {
		return i, errors.New("Parse error with " + string(j))
	}

	return i, err
}
