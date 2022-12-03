package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	ROCK     Shape = 1
	PAPER          = 2
	SCISSORS       = 3
)

const (
	LOST State = 0
	DRAW       = 3
	WIN        = 6
)

type Shape int
type State int

type Round struct {
	firstHand  Shape
	secondHand Shape
}

func toTuple(a []string) (string, string) {
	return a[0], a[1]
}

func mapTo(s string, m map[string]Shape) Shape {
	return m[s]
}

func mapToStrategy(s string, m map[string]func(Shape) Shape) func(Shape) Shape {
	return m[s]
}

func RegularStrategy(oRaw string, uRaw string) int {

	var opponentMap map[string]Shape = map[string]Shape{"A": ROCK, "B": PAPER, "C": SCISSORS}
	var playerMap map[string]Shape = map[string]Shape{"X": ROCK, "Y": PAPER, "Z": SCISSORS}

	opponent := mapTo(oRaw, opponentMap)
	user := mapTo(uRaw, playerMap)

	return score(opponent, user)
}

func score(opponent Shape, user Shape) int {

	if opponent == user {
		return 3 + int(user)
	}

	if user == PAPER {
		if opponent == SCISSORS {
			return 0 + int(user)
		} else if opponent == ROCK {
			return 6 + int(user)
		}
	} else if user == ROCK {
		if opponent == PAPER {
			return 0 + int(user)
		} else if opponent == SCISSORS {
			return 6 + int(user)
		}
	} else if user == SCISSORS {
		if opponent == ROCK {
			return 0 + int(user)
		} else if opponent == PAPER {
			return 6 + int(user)
		}
	}

	return 0 + int(user)
}

func getWinnerShape(s Shape) Shape {
	if s == ROCK {
		return PAPER
	} else if s == PAPER {
		return SCISSORS
	} else {
		return ROCK
	}
}

func getLoserShape(s Shape) Shape {
	if s == ROCK {
		return SCISSORS
	} else if s == PAPER {
		return ROCK
	} else {
		return PAPER
	}
}

func getDrawShape(s Shape) Shape {
	return s
}

func PlayerSelectionStrategy(oRaw string, uRaw string) int {

	var opponentMap map[string]Shape = map[string]Shape{"A": ROCK, "B": PAPER, "C": SCISSORS}
	var userMap map[string]func(Shape) Shape = map[string]func(Shape) Shape{"X": getLoserShape, "Y": getDrawShape, "Z": getWinnerShape}

	opponent := mapTo(oRaw, opponentMap)
	user := mapToStrategy(uRaw, userMap)(opponent)

	return score(opponent, user)
}

func matchResult(path string, calc func(string, string) int) int {

	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var rawMatches []string = strings.Split(string(data), "\n")
	var total int = 0
	for _, pair := range rawMatches {
		oRaw, uRaw := toTuple(strings.Split(pair, " "))
		score := calc(oRaw, uRaw)
		total += score
	}

	return total
}

func main() {
	result1 := matchResult("data/2/input.txt", RegularStrategy)
	result2 := matchResult("data/2/input.txt", PlayerSelectionStrategy)
	fmt.Printf("Part 1 Result %d \n", result1)
	fmt.Printf("Part 2 Result %d \n", result2)
}
