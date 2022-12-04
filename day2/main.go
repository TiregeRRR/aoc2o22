package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	enemyRock      = "A"
	enemyPaper     = "B"
	enemyScissors  = "C"
	playerRock     = "X"
	playerPaper    = "Y"
	playerScissors = "Z"
	playerLose     = "X"
	playerDraw     = "Y"
	playerWin      = "Z"
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
		slc := strings.Split(line, " ")
		score += calcValue1(slc[0], slc[1])
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
		slc := strings.Split(line, " ")
		score += calcValue2(slc[0], slc[1])
	}
	fmt.Printf("Part 2 answer for %s is: %v\n", name, score)
	return nil
}

func calcValue1(enemyTurn, playerTurn string) int {
	switch enemyTurn {
	case enemyRock:
		switch playerTurn {
		case playerRock:
			return 3 + 1
		case playerPaper:
			return 6 + 2
		case playerScissors:
			return 0 + 3
		}
	case enemyPaper:
		switch playerTurn {
		case playerRock:
			return 0 + 1
		case playerPaper:
			return 3 + 2
		case playerScissors:
			return 6 + 3
		}
	case enemyScissors:
		switch playerTurn {
		case playerRock:
			return 6 + 1
		case playerPaper:
			return 0 + 2
		case playerScissors:
			return 3 + 3
		}
	}
	return 0
}

func calcValue2(enemyTurn, playerTurn string) int {
	switch enemyTurn {
	case enemyRock:
		switch playerTurn {
		case playerLose:
			return calcValue1(enemyTurn, playerScissors)
		case playerDraw:
			return calcValue1(enemyTurn, playerRock)
		case playerWin:
			return calcValue1(enemyTurn, playerPaper)
		}
	case enemyPaper:
		switch playerTurn {
		case playerLose:
			return calcValue1(enemyTurn, playerRock)
		case playerDraw:
			return calcValue1(enemyTurn, playerPaper)
		case playerWin:
			return calcValue1(enemyTurn, playerScissors)
		}
	case enemyScissors:
		switch playerTurn {
		case playerLose:
			return calcValue1(enemyTurn, playerPaper)
		case playerDraw:
			return calcValue1(enemyTurn, playerScissors)
		case playerWin:
			return calcValue1(enemyTurn, playerRock)
		}
	}
	return 0
}
