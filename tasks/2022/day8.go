package tasks2022

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"strconv"
	"strings"
)

type Day8 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2022, 8), func() tasks.DailyTask { return Day8{} })
}

func (m Day8) Task1(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	count := 0
	for i := 0; i < len(lines); i++ {
		for n := 0; n < len(lines[i]); n++ {
			line := lines[i]
			value, _ := strconv.Atoi(string(rune(line[n])))
			leftVal := getLeftMax(line, n)
			rightVal := getRightMax(line, n)
			topValue := getTopMax(lines, i, n)
			botValue := getBotMax(lines, i, n)
			if value > leftVal || value > rightVal || value > topValue || value > botValue {
				count++
			}
		}
	}

	result := strconv.Itoa(count)
	return &result, nil
}

func getTopMax(lines []string, rowIndex int, columnIndex int) int {
	maxVal := -1
	for i := 0; i < rowIndex; i++ {
		line := lines[i]
		value, _ := strconv.Atoi(string(rune(line[columnIndex])))
		if value > maxVal {
			maxVal = value

		}
	}

	return maxVal
}

func getBotMax(lines []string, rowIndex int, columnIndex int) int {
	maxVal := -1

	for i := rowIndex + 1; i < len(lines); i++ {
		line := lines[i]
		value, _ := strconv.Atoi(string(rune(line[columnIndex])))
		if value > maxVal {
			maxVal = value
		}
	}
	return maxVal
}

func getRightMax(line string, letterIndex int) int {
	maxVal := -1

	for i := letterIndex + 1; i < len(line); i++ {
		value, _ := strconv.Atoi(string(rune(line[i])))
		if value > maxVal {
			maxVal = value
		}
	}
	return maxVal
}

func getLeftMax(line string, letterIndex int) int {
	maxVal := -1

	for i := 0; i < letterIndex; i++ {
		value, _ := strconv.Atoi(string(rune(line[i])))
		if value > maxVal {
			maxVal = value
		}
	}
	return maxVal
}

func (m Day8) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")

	/*lines = []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}*/

	maxProduct := 0
	for i := 0; i < len(lines); i++ {
		for n := 0; n < len(lines[i]); n++ {
			line := lines[i]
			value, _ := strconv.Atoi(string(rune(line[n])))
			leftNum := getLeftMaxFromValue(line, value, n)
			rightNum := getRightMaxFromValue(line, value, n)
			topNum := getTopMaxFromValue(lines, value, i, n)
			botNum := getBotMaxFromValue(lines, value, i, n)
			product := leftNum * rightNum * topNum * botNum
			if product > maxProduct {
				maxProduct = product
			}
			//fmt.Printf("(%d,%d)=%d : %d [%d,%d,%d,%d]\n", n, i, value, product, leftNum, topNum, rightNum, botNum)

		}
		//fmt.Println()
	}

	result := strconv.Itoa(maxProduct)
	return &result, nil
}

func getTopMaxFromValue(lines []string, compValue int, rowIndex int, columnIndex int) int {
	count := 0
	for i := rowIndex - 1; i >= 0; i-- {
		line := lines[i]
		value, _ := strconv.Atoi(string(rune(line[columnIndex])))
		count++
		if compValue <= value {
			break
		}
	}
	return count
}

func getBotMaxFromValue(lines []string, compValue int, rowIndex int, columnIndex int) int {
	count := 0

	for i := rowIndex + 1; i < len(lines); i++ {
		line := lines[i]
		value, _ := strconv.Atoi(string(rune(line[columnIndex])))
		count++
		if compValue <= value {
			break
		}
	}
	return count
}

func getRightMaxFromValue(line string, compValue int, letterIndex int) int {
	count := 0

	for i := letterIndex + 1; i < len(line); i++ {
		value, _ := strconv.Atoi(string(rune(line[i])))
		count++
		if compValue <= value {
			break
		}
	}
	return count
}

func getLeftMaxFromValue(line string, compValue int, letterIndex int) int {
	count := 0

	for i := letterIndex - 1; i >= 0; i-- {
		value, _ := strconv.Atoi(string(rune(line[i])))
		count++
		if compValue <= value {
			break
		}
	}
	return count
}
