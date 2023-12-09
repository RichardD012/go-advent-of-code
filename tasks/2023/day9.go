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
	//lines = m.getTestData()
	lastNumberSum := 0
	for _, line := range lines {
		numList := convertToNumList(line)
		//lastDifference := calculateLastDifference(numList)
		//lastNumberSum += numList[len(numList)-1] + lastDifference
		diffList := generateDifferenceList(numList)
		lastNumberSum += numList[len(numList)-1] + diffList[len(diffList)-1]
	}
	result := strconv.Itoa(lastNumberSum)
	return &result, nil
}

func (m Day9) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	//lines = m.getTestData()
	lastNumberSum := 0
	for _, line := range lines {
		numList := convertToNumList(line)
		//lastDifference := calculateFirstDifference(numList)
		//lastNumberSum += numList[0] - lastDifference
		diffList := generateDifferenceList(numList)
		lastNumberSum += numList[0] - diffList[0]
	}
	result := strconv.Itoa(lastNumberSum)
	return &result, nil
}

func generateDifferenceList(list []int) []int {
	recurse := false
	var firstDifference *int
	var diffList []int
	for index, intVal := range list[1:] {
		difference := intVal - list[index]
		diffList = append(diffList, difference)
		if firstDifference == nil {
			firstDifference = &difference
		} else {
			if difference != *firstDifference {
				recurse = true
			}
		}
	}
	if recurse == false {
		return diffList
	}
	newList := generateDifferenceList(diffList)
	//prepend the previous diff
	returnList := append([]int{diffList[0] - newList[0]}, diffList[:]...)
	//append the next diff
	returnList = append(returnList, diffList[len(diffList)-1]+newList[len(newList)-1])
	return returnList
}

// Single function
func calculateLastDifference(list []int) int {
	set := false
	recurse := false
	firstDifference := -1
	var diffList []int
	for index, intVal := range list[1:] {
		difference := intVal - list[index]
		diffList = append(diffList, difference)
		if set == false {
			firstDifference = difference
			set = true
		} else {
			if difference != firstDifference {
				recurse = true
			}
		}
	}
	if recurse == false {
		return firstDifference
	}
	return diffList[len(diffList)-1] + calculateLastDifference(diffList)
}

// Single function
func calculateFirstDifference(list []int) int {
	set := false
	recurse := false
	firstDifference := -1
	var diffList []int
	for index, intVal := range list[1:] {
		difference := intVal - list[index]
		diffList = append(diffList, difference)
		if set == false {
			firstDifference = difference
			set = true
		} else {
			if difference != firstDifference {
				recurse = true
			}
		}
	}
	if recurse == false {
		return firstDifference
	}
	return diffList[0] - calculateFirstDifference(diffList)
}

func convertToNumList(line string) []int {
	var returnList []int
	for _, numString := range strings.Fields(line) {
		num, _ := strconv.Atoi(numString)
		returnList = append(returnList, num)
	}
	return returnList
}

func (m Day9) getTestData() []string {
	return []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
}
