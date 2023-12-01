package tasks2022

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"strconv"
)

type Day6 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2022, 6), func() tasks.DailyTask { return Day6{} })
}

func (m Day6) Task1(data *string) (*string, error) {
	str := *data
	resultIndex := 0
	for i := 0; i < len(str)-4; i++ {
		a := string(str[i])
		b := string(str[i+1])
		c := string(str[i+2])
		d := string(str[i+3])
		if a != b && a != c && a != d && b != c && b != d && c != d {
			resultIndex = i + 3
			break
		}
	}
	//it's zero indexed so the 0th index would be the first character
	result := strconv.Itoa(resultIndex + 1)
	return &result, nil
}

func (m Day6) Task2(data *string) (*string, error) {
	str := *data
	resultIndex := 0
	for i := 0; i < len(str)-14; i++ {
		tempWord := str[i : i+14]
		wordHistogram := histogram(tempWord)
		if len(wordHistogram) == 14 {
			resultIndex = i + 13
			break
		}
	}
	//it's zero indexed so the 0th index would be the first character
	result := strconv.Itoa(resultIndex + 1)
	return &result, nil
}

func histogram(s string) map[rune]int {
	hist := make(map[rune]int)
	for _, ch := range s {
		hist[ch]++
	}
	return hist
}
