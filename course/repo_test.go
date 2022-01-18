package main

import (
	"testing"
)

func TestTable(t *testing.T) {
	casesTestRequestСoncatenation := []struct {
		A      []string
		result string
	}{
		{
			A:      []string{"fasdfa", "fsadgfa"},
			result: " ( fasdfa) VALUES (fsadgfa) ",
		},
		{
			A:      []string{"fasdfa", "fsadgfa", "fdafadfa", "vdahgsdrfg"},
			result: " ( fasdfa, fsadgfa) VALUES (fdafadfa, vdahgsdrfg) ",
		},
	}

	for _, testCase := range casesTestRequestСoncatenation {
		res := requestСoncatenation(testCase.A)
		if res != testCase.result {
			t.Errorf("хуйня какая-то")
		}
	}
}
