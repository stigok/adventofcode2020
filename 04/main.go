package main

import "bufio"
import "fmt"
import "os"
import "bytes"
import "regexp"
import "strconv"
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
	fmt.Println(Solve2(lines))
}

var validators map[string]func(string) bool = map[string]func(string) bool{
	"byr": func(s string) bool {
		n, err := strconv.Atoi(s)
		return err == nil && n >= 1920 && n <= 2002
	},
	"iyr": func(s string) bool {
		n, err := strconv.Atoi(s)
		return err == nil && n >= 2010 && n <= 2020
	},
	"eyr": func(s string) bool {
		n, err := strconv.Atoi(s)
		return err == nil && n >= 2020 && n <= 2030
	},
	"hgt": func(s string) bool {
		pat := regexp.MustCompile(`(\d+)(cm|in)`)
		match := pat.FindStringSubmatch(s)
		if match == nil {
			return false
		}
		n, _ := strconv.Atoi(match[1])
		if match[2] == "cm" {
			return n >= 150 && n <= 193
		} else {
			return n >= 59 && n <= 76
		}
	},
	"hcl": func(s string) bool {
		pat := regexp.MustCompile(`#[a-f0-9]{6}`)
		return pat.MatchString(s)
	},
	"ecl": func(s string) bool {
		pat := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
		return pat.MatchString(s)
	},
	"pid": func(s string) bool {
		pat := regexp.MustCompile(`^\d{9}$`)
		return pat.MatchString(s)
	},
}

func Solve1(lines []string) int {
	validCount := 0
	for _, line := range lines {
		valid := true
		for key, _ := range validators {
			if !strings.Contains(line, fmt.Sprintf("%s:", key)) {
				valid = false
				break
			}
		}
		if valid {
			validCount++
		}
	}

	return validCount
}

func Solve2(lines []string) int {
	validCount := 0
	for _, line := range lines {
		fields := make(map[string]string)
		for _, pair := range strings.Split(line, " ") {
			split := strings.Split(pair, ":")
			fields[split[0]] = split[1]
		}

		valid := true
		for key, validate := range validators {
			s, ok := fields[key]
			if !ok || !validate(s) {
				valid = false
				break
			}
		}
		if valid {
			validCount++
		}
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
