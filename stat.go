package main

import "encoding/json"

type stat struct {
	doorCnt           int
	imgCnt            int
	alarmCnt          int
	avgProcessingTime int
}

func encodeStat(s stat) []byte {
	j, _ := json.Marshal(s)
	return j
}

var emptyStat = stat{0, 0, 0, 0}
