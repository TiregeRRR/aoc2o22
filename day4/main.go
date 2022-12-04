package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		slc := strings.Split(line, ",")
		b, err := getContains(slc[0], slc[1])
		if err != nil {
			return err
		}
		if b {
			score++
		}
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
	score := 0
	for buf.Scan() {
		line := buf.Text()
		if line == "" {
			continue
		}
		slc := strings.Split(line, ",")
		b, err := getOverlap(slc[0], slc[1])
		if err != nil {
			return err
		}
		if b {
			score++
		}
	}
	fmt.Printf("Part 2 answer for %s is: %v\n", name, score)
	return nil
}

func getContains(elf1, elf2 string) (bool, error) {
	sl1 := strings.Split(elf1, "-")
	a, err := strconv.Atoi(sl1[0])
	if err != nil {
		return false, err
	}
	b, err := strconv.Atoi(sl1[1])
	if err != nil {
		return false, err
	}
	sl2 := strings.Split(elf2, "-")
	x, err := strconv.Atoi(sl2[0])
	if err != nil {
		return false, err
	}
	y, err := strconv.Atoi(sl2[1])
	if err != nil {
		return false, err
	}
	if a <= x && b >= y {
		return true, nil
	}
	if x <= a && y >= b {
		return true, nil
	}
	return false, nil
}

func getOverlap(elf1, elf2 string) (bool, error) {
	sl1 := strings.Split(elf1, "-")
	a, err := strconv.Atoi(sl1[0])
	if err != nil {
		return false, err
	}
	b, err := strconv.Atoi(sl1[1])
	if err != nil {
		return false, err
	}
	sl2 := strings.Split(elf2, "-")
	x, err := strconv.Atoi(sl2[0])
	if err != nil {
		return false, err
	}
	y, err := strconv.Atoi(sl2[1])
	if err != nil {
		return false, err
	}
	if a <= y && b >= x {
		return true, nil
	}
	return false, nil
}
