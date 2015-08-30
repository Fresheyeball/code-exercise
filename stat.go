package main

import "encoding/json"

type stat struct {
	doorCnt           int
	ImgCnt            int
	alarmCnt          int
	avgProcessingTime int
}

func encodeStat(s stat) []byte {
	j, _ := json.Marshal(s)
	return j
}
