package tasks2023

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"strconv"
	"strings"
)

type Day11 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2023, 11), func() tasks.DailyTask { return Day11{} })
}

func (m Day11) Task1(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	result := strconv.Itoa(len(lines))
	return &result, nil
}

func (m Day11) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	result := strconv.Itoa(len(lines))
	return &result, nil
}

func (m Day11) getTestData() []string {
	return []string{}
}
