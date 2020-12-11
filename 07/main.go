package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println(Solve1(lines))
	fmt.Println(Solve2(lines))
	fmt.Println("already guessed 447")
}

var colorPat *regexp.Regexp = regexp.MustCompile(`^(\w+ \w+) bags? contain`)
var qtyPat *regexp.Regexp = regexp.MustCompile(`(\d+) (\w+ \w+) bag`)

func GetBagSequences(lines []string) ([][]string, Bags) {
	var colors [][]string
	var bags Bags

	for _, s := range lines {
		match := colorPat.FindAllStringSubmatch(s, -1)
		var seq []string
		for _, m := range match {
			seq = append(seq, m[1])
		}
		colors = append(colors, seq)

		bag := &Bag{Color: match[0][1]}
		bags = append(bags, bag)
	}

	// Fill in children
	for i, s := range lines {
		match := qtyPat.FindAllStringSubmatch(s, -1)
		for _, m := range match {
			qty, _ := strconv.Atoi(m[1])
			color := m[2]
			bags[i].Contains = append(bags[i].Contains, Container{
				Bag:   bags.Get(color),
				Count: qty,
			})
		}
	}

	return colors, bags
}

func IndexOf(slice []string, s string) int {
	for i, el := range slice {
		if el == s {
			return i
		}
	}
	return -1
}

func FilterBags(bags []*Bag, f func(*Bag) bool) []*Bag {
	var filtered []*Bag
	for _, b := range bags {
		if f(b) {
			filtered = append(filtered, b)
		}
	}
	return filtered
}

type Bags []*Bag

func (bags Bags) IndexOf(b *Bag) int {
	for i, el := range bags {
		if el.Color == b.Color {
			return i
		}
	}
	return -1
}

func (bags Bags) Get(color string) *Bag {
	for _, b := range bags {
		if b.Color == color {
			return b
		}
	}
	return nil
}

func (all Bags) FindNestedParents(target *Bag) Bags {
	var parents Bags
	var queue Bags
	var e *Bag

	// Initial queue
	queue = make(Bags, len(target.Parents))
	copy(queue, target.Parents)

	for true {
		if len(queue) == 0 {
			break
		}

		// Shift
		e, queue = queue[0], queue[1:]

		if parents.IndexOf(e) >= 0 {
			fmt.Printf("Skipping %s (%p), already exist\n", e.Color, e)
			continue
		}

		fmt.Printf("Found parent %s (%p)\n", e.Color, e)
		parents = append(parents, e)

		// Add the parents themselves to the queue
		for _, p := range e.Parents {
			if parents.IndexOf(p) == -1 {
				queue = append(queue, p)
			}
		}
	}

	return parents
}

type Bag struct {
	Color    string
	Children Bags
	Parents  Bags
	Contains []Container
}

type Container struct {
	Bag   *Bag
	Count int
}

func ParseBags(seqs [][]string) Bags {
	var bags Bags

	// Parse all bags
	for _, seq := range seqs {
		bag := &Bag{
			Color: seq[0],
		}

		bags = append(bags, bag)
	}

	// Populate parents
	for _, bag := range bags {
		for _, seq := range seqs {
			if seq[0] == bag.Color {
				continue
			}
			if IndexOf(seq, bag.Color) > 0 {
				bag.Parents = append(bag.Parents, bags.Get(seq[0]))
			}
		}
	}

	return bags
}

func Solve1(lines []string) int {
	seqs, _ := GetBagSequences(lines)
	bags := ParseBags(seqs)
	mybag := bags.Get("shiny gold")
	parents := bags.FindNestedParents(mybag)
	return len(parents)
}

func Solve2(lines []string) int {
	_, bags := GetBagSequences(lines)
	mybag := bags.Get("shiny gold")

	var c Container
	q := make([]Container, len(mybag.Contains))
	copy(q, mybag.Contains)

	sum := 0
	for true {
		if len(q) == 0 {
			break
		}

		// Array shift
		c, q = q[0], q[1:]
		sum += c.Count

		for i := 0; i < c.Count; i++ {
			q = append(q, c.Bag.Contains...)
		}
	}

	return sum
}
