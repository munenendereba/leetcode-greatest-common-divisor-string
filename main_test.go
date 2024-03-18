package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	string1 string
	string2 string
	divisor string
}{
	{"ABCABC", "ABC", "ABC"},
	{"ABABAB", "ABAB", "AB"},
	{"AABABAB", "ABAB", ""},
	{"AABABAABAB", "AABAB", "AABAB"},
	{"LEET", "CODE", ""},
	{"AAAAAAAAAAAA", "A", "A"},
	{"TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXX"},
}

func TestGcdOfStrings(t *testing.T) {
	for i, test := range tests {
		testname := fmt.Sprintf("running test %d", i)

		t.Run(testname, func(t *testing.T) {
			div := gcdOfStrings(test.string1, test.string2)

			if div != test.divisor {
				t.Errorf("gcdOfStrings(%s,%s) expected %s but got %s", test.string1, test.string2, test.divisor, div)
			}
		})
	}

}
