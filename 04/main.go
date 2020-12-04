package main

import "bufio"
import "fmt"
import "os"
import "bytes"
import "strings"

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(Scan2Lines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println(Solve1(lines))
}

func Solve1(lines []string) int {
	validCount := 0
	req_keys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, line := range lines {
		valid := true
		for _, k := range req_keys {
			if !strings.Contains(line, fmt.Sprintf("%s:", k)) {
				valid = false
				break
			}
		}
		if valid {
			validCount++
		}
		fmt.Printf("%v - %s\n", valid, line)
	}

	return validCount
}

// Modified version of Go's builtin bufio.ScanLines to return strings separated by
// two newlines (instead of one). Returns a string without newlines in it, and trims
// spaces from start and end.
// Does not remove terminal \r's from tokens (dropCR() voided).
// https://github.com/golang/go/blob/master/src/bufio/scan.go#L344-L364
func Scan2Lines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		s := bytes.ReplaceAll(data[0:i+1], []byte("\n"), []byte(" "))
		s = bytes.Trim(s, "\n ")
		return i + 2, s, nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		s := bytes.ReplaceAll(data, []byte("\n"), []byte(" "))
		s = bytes.Trim(s, "\n ")
		return len(data), s, nil
	}
	// Request more data.
	return 0, nil, nil
}
