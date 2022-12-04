package main

import (
	"bufio"
	"fmt"
	"os"
)

var charToInt = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
	'm': 13,
	'n': 14,
	'o': 15,
	'p': 16,
	'q': 17,
	'r': 18,
	's': 19,
	't': 20,
	'u': 21,
	'v': 22,
	'w': 23,
	'x': 24,
	'y': 25,
	'z': 26,
	'A': 27,
	'B': 28,
	'C': 29,
	'D': 30,
	'E': 31,
	'F': 32,
	'G': 33,
	'H': 34,
	'I': 35,
	'J': 36,
	'K': 37,
	'L': 38,
	'M': 39,
	'N': 40,
	'O': 41,
	'P': 42,
	'Q': 43,
	'R': 44,
	'S': 45,
	'T': 46,
	'U': 47,
	'V': 48,
	'W': 49,
	'X': 50,
	'Y': 51,
	'Z': 52,
}

func main() {
	if err := solve("test.txt"); err != nil {
		panic(err)
	}
	if err := solve("a.txt"); err != nil {
		panic(err)
	}
}

func solve(name string) error {
	if err := solve1(name); err != nil {
		return err
	}
	return solve2(name)
}

func solve1(name string) error {
	f, err := os.Open(name)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewScanner(f)
	score := 0
	for buf.Scan() {
		line := buf.Text()
		if line == "" {
			continue
		}
		l := len(line)
		score += getSame(line[0:l/2], line[l/2:l])
	}
	fmt.Printf("Part 1 answer for %s is: %v\n", name, score)
	return nil
}

func solve2(name string) error {
	f, err := os.Open(name)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewScanner(f)
	group := []string{}
	score := 0
	for buf.Scan() {
		line := buf.Text()
		if line == "" {
			continue
		}
		group = append(group, line)
		if len(group) == 3 {
			score += getBadge(group)
			group = []string{}
		}
	}
	fmt.Printf("Part 2 answer for %s is: %v\n", name, score)
	return nil
}

func getSame(p1 string, p2 string) int {
	m := map[rune]struct{}{}
	for _, v := range p1 {
		m[v] = struct{}{}
	}
	for _, v := range p2 {
		if _, ok := m[v]; ok {
			return charToInt[v]
		}
	}
	return 0
}

func getBadge(group []string) int {
	m1 := map[rune]struct{}{}
	m2 := map[rune]struct{}{}
	for _, v := range group[0] {
		m1[v] = struct{}{}
	}
	for _, v := range group[1] {
		m2[v] = struct{}{}
	}
	for _, v := range group[2] {
		if _, ok := m1[v]; ok {
			if _, ok := m2[v]; ok {
				return charToInt[v]
			}
		}
	}
	return 0
}
