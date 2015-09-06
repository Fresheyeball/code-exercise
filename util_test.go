package main

import "math/rand"

func choose(min, max int) int {
	return rand.Intn(max-min) + min
}

func getRandomFrom(xs []string) string {
	return xs[choose(0, len(xs))]
}
