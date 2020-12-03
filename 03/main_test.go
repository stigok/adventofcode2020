package main

import "strings"
import "testing"

var testcase string = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func TestSolve1(t *testing.T) {
	r := strings.Split(testcase, "\n")

	trees := Solve1(r)
	if trees != 7 {
		t.Errorf("expected 7, got %d", trees)
	}
}

func TestSolve2(t *testing.T) {
	r := strings.Split(testcase, "\n")

	trees := Solve2(r)
	if trees != 336 {
		t.Errorf("expected 336, got %d", trees)
	}
}
