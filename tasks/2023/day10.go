package tasks2023

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"strconv"
	"strings"
)

type Day10 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2023, 10), func() tasks.DailyTask { return Day10{} })
}

func (m Day10) Task1(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	result := strconv.Itoa(len(lines))
	return &result, nil
}

func (m Day10) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	result := strconv.Itoa(len(lines))
	return &result, nil
}

func (m Day10) getTestData() []string {
	return []string{}
}
