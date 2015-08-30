package main

import "log"

func attempt(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func isError(_ interface{}, err error) bool {
	return err != nil
}

func safeDiv(x int64, y int64) int64 {
	if y == 0 {
		return 0
	}
	return x / y
}

// tuples are only for returns? bullshit mang!
// func attemptWith(message string, (x int, err error)) int {
// 	return x
// }
