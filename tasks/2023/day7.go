package tasks2023

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"strconv"
	"strings"
)

type Day7 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2023, 7), func() tasks.DailyTask { return Day7{} })
}

func (m Day7) Task1(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	result := strconv.Itoa(len(lines))
	return &result, nil
}

func (m Day7) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	result := strconv.Itoa(len(lines))
	return &result, nil
}
func (m Day7) getTestData() []string {
	return []string{}
}
