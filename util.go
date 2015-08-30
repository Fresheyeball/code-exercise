package main

import "log"

func attempt(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// tuples are only for returns? bullshit mang!
// func attemptWith(message string, (x int, err error)) int {
// 	return x
// }
