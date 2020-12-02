package main

import "testing"

func TestVersion1(t *testing.T) {
	testcases := map[string]bool{
		"1-3 a: abcde":     true,
		"1-3 b: cdefg":     false,
		"2-9 c: ccccccccc": true,
	}
	for s, expected := range testcases {
		p, err := ParsePasswordPolicy(s)
		if err != nil {
			t.Error(err)
		}
		if p.IsValidV1() != expected {
			t.Errorf("%v expected to be %v", p, expected)
		}
	}
}

func TestVersion2(t *testing.T) {
	testcases := map[string]bool{
		"1-3 a: abcde":     true,
		"1-3 b: cdefg":     false,
		"2-9 c: ccccccccc": false,
	}
	for s, expected := range testcases {
		p, err := ParsePasswordPolicy(s)
		if err != nil {
			t.Error(err)
		}
		if p.IsValidV2() != expected {
			t.Errorf("%v expected to be %v", p, expected)
		}
	}
}
