package main

import "math/rand"

const checkSize int = 1000

func check(f func()) {
	forN(checkSize, f)
}

func choose(min, max int) int {
	return rand.Intn(max-min) + min
}

func getRandomFrom(xs []string) string {
	return xs[choose(0, len(xs))]
}
