package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println(Solve1(lines))
	fmt.Println(Solve2(lines))
}

func Solve1(iids []string) int {
	var sids []int
	for _, iid := range iids {
		_, _, sid := DecodeItinerary(iid)
		sids = append(sids, sid)
	}
	return MaxInt(sids)
}

func Solve2(iids []string) int {
	var sids []int
	for _, iid := range iids {
		_, _, sid := DecodeItinerary(iid)
		sids = append(sids, sid)
	}

	sort.Ints(sids)

	for i, n := range sids {
		if sids[i+1] == n+2 {
			return n + 1
		}
	}
	return -1
}

// Get ticket information by full itinerary ID
func DecodeItinerary(iid string) (row int, col int, seat int) {
	rowid := iid[0:7]
	colid := iid[7:]
	row = DecodeIidPart(rowid)
	col = DecodeIidPart(colid)
	seat = GetSeatId(row, col)
	return
}

// Determine seat number from row or columns part of itinerary ID
func DecodeIidPart(iid string) int {
	n := 0
	d := len(iid) - 1

	for i, c := range iid {
		if c == 'B' || c == 'R' {
			n ^= 1 << (d - i)
		}
	}

	return n
}

// Return seat ID based on full itinerary ID
func GetSeatId(row, col int) int {
	return row*8 + col
}

// Return the biggest number in the slice
func MaxInt(nums []int) int {
	max := -1
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}
