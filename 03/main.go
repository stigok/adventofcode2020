package main

import "bufio"
import "fmt"
import "io"
import "os"

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	trees := Solve1(file)
	fmt.Println(trees)
}

func Solve1(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	line := scanner.Text()

	x := 0
	d := 3
	n := len(line)
	trees := 0

	for scanner.Scan() {
		fmt.Println(x, d, n, trees)
		x += d
		line = scanner.Text()
		if line[x%n] == '#' {
			trees++
		}
	}

	return trees
}
