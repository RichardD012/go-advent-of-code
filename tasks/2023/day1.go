package tasks2023

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"regexp"
	"strconv"
	"strings"
)

type Day1 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2023, 1), func() tasks.DailyTask { return Day1{} })
}

func (m Day1) Task1(data *string) (*string, error) {
	re, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return nil, err
	}
	sum := 0
	lines := strings.Split(*data, "\n")
	for _, line := range lines {
		// Replace non-digit characters with an empty string
		onlyDigits := re.ReplaceAllString(line, "")
		num, _ := strconv.Atoi(string(onlyDigits[0]) + string(onlyDigits[len(onlyDigits)-1]))
		sum += num
	}

	result := strconv.Itoa(sum)
	return &result, nil
}

func (m Day1) Task2(data *string) (*string, error) {

	re, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return nil, err
	}
	sum := 0
	lines := strings.Split(*data, "\n")
	// Use the test data
	//lines := []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}
	for _, line := range lines {
		line = replaceNumbers(line)
		onlyDigits := re.ReplaceAllString(line, "")
		num, _ := strconv.Atoi(string(onlyDigits[0]) + string(onlyDigits[len(onlyDigits)-1]))
		sum += num
	}
	result := strconv.Itoa(sum)
	return &result, nil
}

// this is whildly inefficient for what its doing - I thought it was a _single_ replacement so that if it was fiveight it could only be 5ight and not 58
func replaceNumbers(line string) string {
	index := 0
	returnString := ""
	for index < len(line) {
		currentString := line[index:]
		if strings.HasPrefix(currentString, "one") {
			returnString += "1"
		} else if strings.HasPrefix(currentString, "two") {
			returnString += "2"
		} else if strings.HasPrefix(currentString, "three") {
			returnString += "3"
		} else if strings.HasPrefix(currentString, "four") {
			returnString += "4"
		} else if strings.HasPrefix(currentString, "five") {
			returnString += "5"
		} else if strings.HasPrefix(currentString, "six") {
			returnString += "6"
		} else if strings.HasPrefix(currentString, "seven") {
			returnString += "7"
		} else if strings.HasPrefix(currentString, "eight") {
			returnString += "8"
		} else if strings.HasPrefix(currentString, "nine") {
			returnString += "9"
		} else {
			returnString += string(line[index])
		}
		index++
	}
	return returnString
}
