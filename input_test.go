package main

import (
	"strings"
	"testing"

	"github.com/google/gofuzz"
)

func TestDecode(t *testing.T) {
	var s string
	fuzzy := fuzz.New()
	forN(100, func() {
		fuzzy.Fuzz(&s)
		kind := s
		data := "{\"Type\":\"" + kind + "\"}"
		a, err := decode([]byte(data))
		expected := input{kind}

		// bail! diminishing returns on handling json parsing of backslash
		if strings.Contains(kind, `\`) {
			return
		}

		isValid := kind != "" && !strings.Contains(kind, `"`)

		if !isValid && err == nil {
			t.Fatal("decode should have failed, but did'nt with ", kind, " and ", a, err)
		}

		if isValid && err != nil && a != expected {
			t.Fatal("decode did not parse as expected with ", kind, " and ", a)
		}

	})
}
