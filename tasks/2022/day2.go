package tasks2022

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"strconv"
	"strings"
)

type Day2 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2022, 2), func() tasks.DailyTask { return Day2{} })
}

type Hand int

const (
	Rock Hand = iota
	Paper
	Scissors
)

func (m Day2) Task1(data *string) (*string, error) {
	sum := 0
	lines := strings.Split(*data, "\n")
	for _, line := range lines {
		opponent := getOpponent(line[0])
		mine := getMine(line[2])
		lineScore := generateLineScore(opponent, mine)
		sum += lineScore
	}
	result := strconv.Itoa(sum)
	return &result, nil
}

func (m Day2) Task2(data *string) (*string, error) {

	sum := 0
	lines := strings.Split(*data, "\n")
	for _, line := range lines {
		opponent := getOpponent(line[0])
		mine := getMineBasedOnOpponent(line[2], opponent)
		lineScore := generateLineScore(opponent, mine)
		sum += lineScore
	}
	result := strconv.Itoa(sum)
	return &result, nil
}

func generateLineScore(opponent Hand, mine Hand) int {
	lineScore := 0
	if opponent == mine {
		lineScore += 3 //draw
	} else if (mine == Rock && opponent == Scissors) || (mine == Scissors && opponent == Paper) || (mine == Paper && opponent == Rock) {
		lineScore += 6
	} else {
		lineScore += 0
	}
	switch mine {
	case Rock:
		lineScore += 1
	case Paper:
		lineScore += 2
	case Scissors:
		lineScore += 3
	}
	return lineScore
}

func getMineBasedOnOpponent(outcome byte, opponent Hand) Hand {
	if outcome == 'X' {
		//X = lose
		switch opponent {
		case Scissors:
			return Paper
		case Paper:
			return Rock
		case Rock:
			return Scissors
		}
	} else if outcome == 'Y' {
		//Y = draw
		return opponent
	} else {
		//Z = win
		switch opponent {
		case Scissors:
			return Rock
		case Paper:
			return Scissors
		case Rock:
			return Paper
		}
	}
	return Rock
}

func getOpponent(letter byte) Hand {
	switch letter {
	case 'A', 'a':
		return Rock
	case 'B', 'b':
		return Paper
	case 'C', 'c':
		return Scissors
	}
	return Rock
}

func getMine(letter byte) Hand {
	switch letter {
	case 'X', 'x':
		return Rock
	case 'Y', 'y':
		return Paper
	case 'Z', 'z':
		return Scissors
	}
	return Rock
}
