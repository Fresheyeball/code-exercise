package main

import "log"

type println string

func runPrintln(println println) {
	log.Println(println)
}
