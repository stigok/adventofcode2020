package main

import "bufio"
import "fmt"
import "os"
import "strconv"

func main() {
	var nums []int

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("atoi conversion failed for '%s'", s)
		}
		nums = append(nums, n)
	}

	res, err := Solve1(nums)
	if err != nil {
		fmt.Println("solve1 failed:", err)
	}
	fmt.Println("Answer1:", res)

	res, err = Solve2(nums)
	if err != nil {
		fmt.Println("solve2 failed:", err)
	}
	fmt.Println("Answer2:", res)

}

// Given an int slice, determine which two numbers adds up to 2020 and return
// the product of them.
func Solve1(nums []int) (int, error) {
	for i, _ := range nums {
		for j, _ := range nums {
			fmt.Println(i, j)
			if i == j {
				continue
			}
			if nums[i]+nums[j] == 2020 {
				return nums[i] * nums[j], nil
			}
		}
	}
	return -1, fmt.Errorf("no answer found")
}

// Given an int slice, determine which three numbers adds up to 2020 and return
// the product of them.
func Solve2(nums []int) (int, error) {
	for i, a := range nums {
		for j, b := range nums {
			for k, c := range nums {
				if i == j || i == k || j == k {
					continue
				}
				if a+b+c == 2020 {
					return a * b * c, nil
				}
			}
		}
	}
	return -1, fmt.Errorf("no answer found")
}
