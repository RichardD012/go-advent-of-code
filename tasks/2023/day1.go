package tasks2023

import (
	"fmt"
	"github.com/RichardD012/go-advent-of-code/tasks"
	"github.com/RichardD012/go-advent-of-code/utils"
)

type Day1 struct {
}

func init() {
	fmt.Println("registering")
	tasks.RegisterStruct(tasks.TaskKey(2023, 1), func() tasks.DailyTask { return Day1{} })
}

func (m Day1) Task1(data *string) (*string, error) {
	_, err := utils.StringToIntSlice(data)
	if err != nil {
		return nil, err
	}
	result := fmt.Sprint("Task 1 Result")
	return &result, nil
}

func (m Day1) Task2(*string) (*string, error) {
	result := fmt.Sprint("Task 2 Result")
	return &result, nil
}
