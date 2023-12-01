package tasks2022

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"sort"
	"strconv"
	"strings"
)

type Day1 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2022, 1), func() tasks.DailyTask { return Day1{} })
}

func (m Day1) Task1(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	maxCalories := 0
	currentCalories := 0
	for _, line := range lines {
		if line == "" {
			//fmt.Printf("Max: %d - Current: %d\n\n", maxCalories, currentCalories)
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
		} else {
			//fmt.Printf("%s\n", line)
			cals, _ := strconv.Atoi(line)
			currentCalories += cals
		}
	}
	if currentCalories > maxCalories {
		maxCalories = currentCalories
	}
	result := strconv.Itoa(maxCalories)
	return &result, nil
}

func (m Day1) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	var resultList []int
	currentCalories := 0
	for _, line := range lines {
		if line == "" {
			resultList = append(resultList, currentCalories)
			currentCalories = 0
		} else {
			cals, _ := strconv.Atoi(line)
			currentCalories += cals
		}
	}
	resultList = append(resultList, currentCalories)
	sort.Ints(resultList)
	sum := 0
	for _, value := range resultList[len(resultList)-3:] {
		sum += value
	}
	result := strconv.Itoa(sum)
	return &result, nil
}
