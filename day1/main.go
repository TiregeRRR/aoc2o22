package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var errEOF = errors.New("EOF")

func main() {
	err := solve("test.txt")
	if err != nil {
		panic(err)
	}
	err = solve("a.txt")
	if err != nil {
		panic(err)
	}
}

func solve(name string) error {
	if err := solvePart1(name); err != nil {
		return err
	}
	return solvePart2(name)
}

func solvePart1(name string) error {
	f, err := os.Open(name)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewScanner(f)
	max := 0
	curElf := 0
	for buf.Scan() {
		line := buf.Text()
		if line == "" {
			if curElf > max {
				max = curElf
			}
			curElf = 0
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		curElf += n
	}
	if curElf > max {
		max = curElf
	}
	fmt.Printf("Part 1 answer for %s is: %v\n", name, max)
	return nil
}

func solvePart2(name string) error {
	f, err := os.Open(name)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewScanner(f)
	max := make([]int, 3)
	curElf := 0
	for buf.Scan() {
		line := buf.Text()
		if line == "" {
			max = shiftArr(max, curElf)
			curElf = 0
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		curElf += n
	}
	max = shiftArr(max, curElf)
	fmt.Printf("Part 2 answer for %s is: %v\n", name, sum(max))
	return nil
}

func shiftArr(max []int, curElf int) []int {
	for i := range max {
		if curElf > max[i] {
			switch i {
			case 0:
				max[2] = max[1]
				max[1] = max[0]
				max[0] = curElf
				return max
			case 1:
				max[2] = max[1]
				max[1] = curElf
				return max
			case 2:
				max[2] = curElf
				return max
			}
		}
	}
	return max
}

func sum(max []int) int {
	sum := 0
	for i := range max {
		sum += max[i]
	}
	return sum
}
