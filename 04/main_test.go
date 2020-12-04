package main

import "bufio"
import "strings"
import "testing"

var testcase string = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

func TestSolve1(t *testing.T) {
	els := strings.Split(testcase, "\n\n")

	num := Solve1(els)
	if num != 2 {
		t.Errorf("expected 2, got %d", num)
	}
}

func TestScan2Lines(t *testing.T) {
	s := `hello my
world
!

foo
bar baz
broken

baz`
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(Scan2Lines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if lines[0] != "hello my world !" {
		t.Errorf("wanted 'hello world !', got '%s'", lines[0])
	}
	if lines[1] != "foo bar baz broken" {
		t.Errorf("wanted 'foo bar', got '%s'", lines[1])
	}
	if lines[2] != "baz" {
		t.Errorf("wanted 'baz', got '%s'", lines[2])
	}
}
