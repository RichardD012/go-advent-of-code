package tasks2023

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"strconv"
	"strings"
)

type Day6 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2023, 6), func() tasks.DailyTask { return Day6{} })
}

func (m Day6) Task1(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	lines = m.getTestData()
	result := strconv.Itoa(len(lines))
	return &result, nil
}

func (m Day6) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	result := strconv.Itoa(len(lines))
	return &result, nil
}

func (m Day6) getTestData() []string {
	return []string{}
}
