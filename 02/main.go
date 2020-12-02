package main

import "bufio"
import "fmt"
import "os"
import "regexp"
import "strconv"
import "strings"

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	validPasswordsV1 := 0
	validPasswordsV2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		p, err := ParsePasswordPolicy(s)
		if err != nil {
			fmt.Println(err)
		}
		if p.IsValidV1() {
			validPasswordsV1++
		}
		if p.IsValidV2() {
			validPasswordsV2++
		}
	}

	fmt.Printf("found %d (v1) and %d (v2) valid passwords\n", validPasswordsV1, validPasswordsV2)
}

type PasswordPolicy struct {
	MinOccurences int
	MaxOccurences int
	Char          string
	Password      string
}

func (p PasswordPolicy) IsValidV1() bool {
	n := strings.Count(p.Password, p.Char)
	return n >= p.MinOccurences && n <= p.MaxOccurences
}

func (p PasswordPolicy) IsValidV2() bool {
	a := p.Password[p.MinOccurences-1] == p.Char[0]
	b := p.Password[p.MaxOccurences-1] == p.Char[0]
	return (a && !b) || (!a && b)
}

func ParsePasswordPolicy(s string) (pol PasswordPolicy, err error) {
	pat := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
	match := pat.FindStringSubmatch(s)
	if match == nil {
		return pol, fmt.Errorf("invalid policy format")
	}
	pol.MinOccurences, _ = strconv.Atoi(match[1])
	pol.MaxOccurences, _ = strconv.Atoi(match[2])
	pol.Char = match[3]
	pol.Password = match[4]
	return pol, nil
}
