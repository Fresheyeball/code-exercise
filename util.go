package main

import "log"

func attempt(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func attemptWith(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func attemptGet(x interface{}, err error) interface{} {
	attempt(err)
	return x
}

func isError(_ interface{}, err error) bool {
	return err == nil
}

func safeDiv(x int64, y int64) int64 {
	if y == 0 {
		return 0
	}
	return x / y
}
