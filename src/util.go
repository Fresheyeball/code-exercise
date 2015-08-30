package main

import "log"

func fatality(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
