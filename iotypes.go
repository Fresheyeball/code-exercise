package main

import (
	"errors"
	"io/ioutil"
	"log"
)

type println string

func runPrintln(println println) {
	if println != "" {
		log.Println(println)
	}
}

type readFile string

func runReadfile(readFile readFile) ([]byte, error) {
	if readFile != "" {
		return ioutil.ReadFile(string(readFile))
	}
	return []byte{}, errors.New("file name cannot be empty")
}
