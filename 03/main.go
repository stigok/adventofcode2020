package main

import "bufio"
import "fmt"
import "os"

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("answer 1", Solve1(lines))
	fmt.Println("answer 2", Solve2(lines))
}

func Solver(lines []string, dx, dy int) int {
	trees := 0

	x := 0
	z := len(lines[0])
	//fmt.Printf("dx %d, dy%d, z=%d\n", dx, dy, z)

	for y := dy; y < len(lines); y += dy {
		x += dx
		if lines[y][x%z] == '#' {
			trees++
		}
		//fmt.Printf("x%d y%d = %d\n", x, y, trees)
	}

	return trees
}

func Solve1(lines []string) int {
	return Solver(lines, 3, 1)
}

func Solve2(lines []string) int {
	answer := 1
	tasks := [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	for _, e := range tasks {
		dx := e[0]
		dy := e[1]
		sum := Solver(lines, dx, dy)
		answer *= sum
		//fmt.Println(dx, dy, sum, answer)
	}
	return answer
}
