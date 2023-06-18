package main

import (
	"fmt"
	"testing"
)

func TestCleanStr(t *testing.T) {
	cases := []struct {
		input  string
		expect string
	}{
		{
			input:  "",
			expect: "",
		},
		{
			input:  "hello world!",
			expect: "hello world!",
		},
		{
			input:  "kerfufflesharbertfornax",
			expect: "kerfufflesharbertfornax",
		},
		{
			input:  "kerfuffle sharbert fornax",
			expect: "**** **** ****",
		},
		{
			input:  "KerFufflE sHarbeRt Fornax",
			expect: "**** **** ****",
		},
		{
			input:  "Kerfuffle!  Sharbert? Fornax. KerFufflE sHarbeRt Fornax",
			expect: "Kerfuffle!  Sharbert? Fornax. **** **** ****",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test Case #%v:", i), func(t *testing.T) {
			cleaned := cleanStr(c.input)
			if cleaned != c.expect {
				t.Errorf("Unexpected: " + cleaned)
				return
			}
		})
	}
}
