package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println(Solve1(lines))
}

var colorPat *regexp.Regexp = regexp.MustCompile(`(\w+ \w+) bag`)

func GetBagSequences(lines []string) [][]string {
	var colors [][]string

	for _, s := range lines {
		match := colorPat.FindAllStringSubmatch(s, -1)
		var seq []string
		for _, m := range match {
			seq = append(seq, m[1])
		}
		colors = append(colors, seq)
	}

	return colors
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
	seqs := GetBagSequences(lines)
	bags := ParseBags(seqs)
	mybag := bags.Get("shiny gold")
	parents := bags.FindNestedParents(mybag)
	return len(parents)
}
