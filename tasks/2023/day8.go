package tasks2023

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"strconv"
	"strings"
)

type Day8 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2023, 8), func() tasks.DailyTask { return Day8{} })
}

func (m Day8) Task1(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	result := strconv.Itoa(len(lines))
	return &result, nil
}

func (m Day8) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	result := strconv.Itoa(len(lines))
	return &result, nil
}

func (m Day8) getTestData() []string {
	return []string{}
}
