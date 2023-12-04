package tasks2023

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"strconv"
	"strings"
)

type Day4 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2023, 4), func() tasks.DailyTask { return Day4{} })
}

func (m Day4) Task1(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	sum := 0
	for _, line := range lines {
		cardValue := 0
		card := strings.Split(line, ": ")
		game := strings.Split(card[1], "|")
		winningSet := game[0]
		myNums := game[1]
		winningMap := make(map[int]int)
		for _, numS := range strings.Fields(winningSet) {
			num, _ := strconv.Atoi(numS)
			winningMap[num] = 1
		}
		for _, numS := range strings.Fields(myNums) {
			num, _ := strconv.Atoi(numS)
			_, exists := winningMap[num]
			if exists {
				if cardValue == 0 {
					cardValue = 1
				} else {
					cardValue *= 2
				}
			}
		}
		sum += cardValue
	}
	result := strconv.Itoa(sum)
	return &result, nil
}

func (m Day4) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	/*lines = []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}*/
	sum := 0
	var lineList []LineStruct
	for _, lineVal := range lines {
		lineList = append(lineList, LineStruct{line: lineVal, count: 1})
	}
	for index, lineStruct := range lineList {
		line := lineStruct.line
		cardValue := 0
		card := strings.Split(line, ": ")
		game := strings.Split(card[1], "|")
		winningSet := game[0]
		myNums := game[1]
		winningMap := make(map[int]int)
		for _, numS := range strings.Fields(winningSet) {
			num, _ := strconv.Atoi(numS)
			winningMap[num] = 1
		}
		for _, numS := range strings.Fields(myNums) {
			num, _ := strconv.Atoi(numS)
			_, exists := winningMap[num]
			if exists {
				cardValue++
			}
		}
		if cardValue > 0 {
			for i := 1; i <= cardValue; i++ {
				lineList[i+index].count += lineStruct.count * 1
			}
		}
	}
	for _, lineStruct := range lineList {
		sum += lineStruct.count
	}
	result := strconv.Itoa(sum)
	return &result, nil
}

type LineStruct struct {
	line  string
	count int
}
