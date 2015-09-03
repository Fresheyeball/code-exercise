package main

import (
	"testing"

	"github.com/google/gofuzz"
)

func fuzzString() string {
	var s string
	fuzz.New().Fuzz(&s)
	return s
}

func TestDecode(t *testing.T) {
	for i := 1; i <= 100; i++ {
		kind := fuzzString()
		data := "{\"Type\":\"" + kind + "\"}"
		a, _ := decode([]byte(data))
		expected := input{kind}
		if a != expected {
			t.Fatal("decode did not parse as expected with", kind)
		}
	}

}
