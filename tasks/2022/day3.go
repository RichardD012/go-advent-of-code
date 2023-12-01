package tasks2022

import (
	"fmt"
	"github.com/RichardD012/go-advent-of-code/tasks"
	"regexp"
	"strconv"
	"strings"
)

type Day3 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2022, 3), func() tasks.DailyTask { return Day3{} })
}

func (m Day3) Task1(data *string) (*string, error) {
	sum := 0
	lines := strings.Split(*data, "\n")
	for _, line := range lines {
		word1 := line[0 : len(line)/2]
		word2 := line[len(line)/2:]
		re, err := regexp.Compile(fmt.Sprintf("[^%s]+", word1))
		if err != nil {
			return nil, err
		}
		onlyWord2 := re.ReplaceAllString(word2, "")
		lineValue := getValue(onlyWord2[0])
		sum += lineValue
	}
	result := strconv.Itoa(sum)
	return &result, nil
}

func (m Day3) Task2(data *string) (*string, error) {
	sum := 0
	lines := strings.Split(*data, "\n")
	for i := 0; i < len(lines); i += 3 {
		line1 := lines[i]
		line2 := lines[i+1]
		line3 := lines[i+2]
		re1, err := regexp.Compile(fmt.Sprintf("[^%s]+", line1))
		if err != nil {
			return nil, err
		}
		onlyLine1 := re1.ReplaceAllString(line2, "")
		re2, err := regexp.Compile(fmt.Sprintf("[^%s]+", onlyLine1))
		if err != nil {
			return nil, err
		}
		onlyAll := re2.ReplaceAllString(line3, "")
		lineValue := getValue(onlyAll[0])
		sum += lineValue
	}
	result := strconv.Itoa(sum)
	return &result, nil
}

func getValue(letter byte) int {
	if letter >= 'a' && letter <= 'z' {
		// 'a' has ASCII value 97, so subtract 96 to make 'a' start at 1
		return int(letter) - 96
	} else if letter >= 'A' && letter <= 'Z' {
		// 'A' has ASCII value 65, so subtract 64 and add 26 to start 'A' at 27
		return int(letter) - 64 + 26
	}
	return 0
}
