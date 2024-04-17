package tasks2023

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"strconv"
	"strings"
)

type Day12 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2023, 12), func() tasks.DailyTask { return Day12{} })
}

func (m Day12) Task1(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	result := strconv.Itoa(len(lines))
	return &result, nil
}

func (m Day12) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	result := strconv.Itoa(len(lines))
	return &result, nil
}

func (m Day12) getTestData() []string {
	return []string{}
}
