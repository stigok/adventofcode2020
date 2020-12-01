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

	res, err := Solve(nums)
	if err != nil {
		fmt.Println("solve failed:", err)
	}

	fmt.Println("Answer:", res)
}

// Given an int slice, determine which two numbers adds up to 2020 and return
// the product of them.
func Solve(nums []int) (int, error) {
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
