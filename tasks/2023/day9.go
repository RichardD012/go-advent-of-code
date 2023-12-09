package tasks2023

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"strconv"
	"strings"
)

type Day9 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2023, 9), func() tasks.DailyTask { return Day9{} })
}

func (m Day9) Task1(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	lastNumberSum := 0
	for _, line := range lines[0:1] {
		numList := convertToNumList(line)
		lastDifference := calculateLastDifference(numList)
		lastNumberSum += numList[len(numList)-1] + lastDifference
	}
	result := strconv.Itoa(lastNumberSum)

	return &result, nil
}

func calculateLastDifference(list []int) int {
	returnInt := 0
	return returnInt
}

func convertToNumList(line string) []int {
	var returnList []int
	for _, numString := range strings.Fields(line) {
		num, _ := strconv.Atoi(numString)
		returnList = append(returnList, num)
	}
	return returnList
}

func (m Day9) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	result := strconv.Itoa(len(lines))
	return &result, nil
}

func (m Day9) getTestData() []string {
	return []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
}
