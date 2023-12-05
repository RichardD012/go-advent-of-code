package tasks2023

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"regexp"
	"strconv"
	"strings"
)

type Day2 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2023, 2), func() tasks.DailyTask { return Day2{} })
}

func (m Day2) Task1(data *string) (*string, error) {
	result := "No Result"
	regexPattern := regexp.MustCompile("[^0-9]")
	lines := strings.Split(*data, "\n")
	/*lines = m.getTestData()*/
	idSum := 0
	for _, line := range lines {

		splitData := strings.Split(line, ":")
		roundNumber, _ := strconv.Atoi(regexPattern.ReplaceAllString(splitData[0], ""))
		throws := strings.Split(splitData[1], ";")
		redPossible := true
		bluePossible := true
		greenPossible := true
		for _, throw := range throws {
			iGroup := strings.Split(throw, ",")
			//determine if each subsection here is possible
			for _, set := range iGroup {
				color, nums := determineOutput(set, *regexPattern)
				switch color {
				case Red:
					if nums > 12 {
						redPossible = false
					}
					break
				case Green:
					if nums > 13 {
						greenPossible = false
					}
					break
				case Blue:
					if nums > 14 {
						bluePossible = false
					}
					break
				}
			}
		}
		//determine if round possible
		if redPossible && bluePossible && greenPossible {
			idSum += roundNumber
		}
	}
	result = strconv.Itoa(idSum)
	return &result, nil
}

func (m Day2) Task2(data *string) (*string, error) {
	result := "No Result"
	regexPattern := regexp.MustCompile("[^0-9]")
	lines := strings.Split(*data, "\n")
	idSum := 0
	for _, line := range lines {

		splitData := strings.Split(line, ":")
		colorCodeGroups := strings.Split(splitData[1], ";")
		minRed, minBlue, minGreen := 0, 0, 0
		for _, group := range colorCodeGroups {
			iGroup := strings.Split(group, ",")
			//determine if each subsection here is possible
			for _, set := range iGroup {
				color, nums := determineOutput(set, *regexPattern)
				switch color {
				case Red:
					if nums >= minRed {
						minRed = nums
					}
					break
				case Green:
					if nums >= minGreen {
						minGreen = nums
					}
					break
				case Blue:
					if nums >= minBlue {
						minBlue = nums
					}
					break
				}
			}
		}
		roundPower := minRed * minBlue * minGreen
		idSum += roundPower

	}
	result = strconv.Itoa(idSum)
	return &result, nil
}

func determineOutput(code string, regexPattern regexp.Regexp) (Color, int) {
	var outputColor Color
	count := 0
	//regexPattern := regexp.MustCompile("[^0-9]")
	if strings.Contains(code, "green") {
		outputColor = Green
	} else if strings.Contains(code, "blue") {
		outputColor = Blue
	} else {
		outputColor = Red
	}
	value := regexPattern.ReplaceAllString(code, "")
	count, _ = strconv.Atoi(value)
	return outputColor, count
}

type Color int

const (
	Red Color = iota
	Green
	Blue
)

func (m Day2) getTestData() []string {
	return []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
}
