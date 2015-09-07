package main

import "log"

func attempt(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func attemptGet(x interface{}, err error) interface{} {
	attempt(err)
	return x
}

func safeDiv(x int64, y int64) int64 {
	if y == 0 {
		return 0
	}
	return x / y
}

func forN(i int, f func()) {
	for j := 1; j <= i; j++ {
		f()
	}
}
