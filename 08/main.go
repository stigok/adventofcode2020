package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("%d lines seen\n", len(lines))
	fmt.Println(Solve1(lines))
}

var acc int
var pc int

func Solve1(instructions []string) int {
	hist := make(map[int]bool)

	ops := make(map[string]func(arg int))
	ops["nop"] = func(n int) {}
	ops["acc"] = func(n int) { acc += n }
	ops["jmp"] = func(n int) { pc += n - 1 }

	for pc = 0; pc < len(instructions); pc++ {
		if _, visited := hist[pc]; visited {
			fmt.Printf("already visited %d, exiting!\n", pc)
			break
		}
		hist[pc] = true

		ins := instructions[pc]
		op := ins[0:3]
		arg, err := strconv.Atoi(ins[4:])
		if err != nil {
			fmt.Printf("failed to convert arg '%v' for op '%s'\n", arg, op)
		}

		fmt.Printf("%d: %s %v (acc %v)\n", pc, ins, arg, acc)
		ops[op](arg)
	}

	return acc
}
