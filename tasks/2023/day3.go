package tasks2023

import (
	"fmt"
	"github.com/RichardD012/go-advent-of-code/tasks"
	"strings"
	"unicode"
)

type Day3 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2023, 3), func() tasks.DailyTask { return Day3{} })
}

func (m Day3) Task1(data *string) (*string, error) {
	result := "No Result"
	lines := strings.Split(*data, "\n")
	lines = []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			char := rune(line[i])
			if unicode.IsDigit(char) {
				fmt.Printf("'%c' is a digit\n", char)
			} else {
				fmt.Printf("'%c' is not a digit\n", char)
			}
		}
	}
	return &result, nil
}

func (m Day3) Task2(data *string) (*string, error) {

	result := "No Result"
	return &result, nil
}
