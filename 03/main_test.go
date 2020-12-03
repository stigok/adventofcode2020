package main

import "strings"
import "testing"

func TestMain(t *testing.T) {
	r := strings.NewReader(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)

	trees := Solve1(r)
	if trees != 7 {
		t.Errorf("expected 7 trees, got %d", trees)
	}
}
