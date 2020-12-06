package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	utils "github.com/stigok/go-utils"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(utils.ScanTwoConsecutiveNewlines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("q1: %d\n", Solve1(lines))
	fmt.Printf("q2: %d\n", Solve2(lines))
}

func Solve1(groups []string) int {
	var sum int
	for _, g := range groups {
		gchars := make(map[byte]struct{})
		for c := byte('a'); c < 'z'+1; c++ {
			if strings.IndexByte(g, c) >= 0 {
				gchars[c] = struct{}{}
			}
		}
		sum += len(gchars)
	}
	return sum
}

func Solve2(groups []string) int {
	var sum int
	for _, g := range groups {
		for c := byte('a'); c < 'z'+1; c++ {
			answers := strings.Split(g, " ")
			if AllContains(answers, c) {
				sum++
			}
		}
	}
	return sum
}

func AllContains(slice []string, b byte) bool {
	for _, s := range slice {
		if strings.IndexByte(s, b) == -1 {
			return false
		}
	}
	return true
}
