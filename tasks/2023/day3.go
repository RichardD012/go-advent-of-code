package tasks2023

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"strconv"
	"strings"
	"unicode"
)

type Day3 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2023, 3), func() tasks.DailyTask { return Day3{} })
}

func (m Day3) Task1(data *string) (*string, error) {
	//result := "No Result"
	lines := strings.Split(*data, "\n")
	sum := 0
	for lineIndex, line := range lines {
		for startIndex := 0; startIndex < len(line); startIndex++ {
			char := rune(line[startIndex])
			//if we have a digit
			if unicode.IsDigit(char) {
				//get the last in a sequence of digits
				endIndex := getEndNum(line, startIndex)
				//convert to a number
				num, _ := strconv.Atoi(line[startIndex:endIndex])
				anySymbol := false
				//check above
				if lineIndex > 0 {
					anySymbol = anySymbol || checkLineForSymbol(lines[lineIndex-1], startIndex, endIndex)
				}
				//check below
				if lineIndex < len(lines)-1 {
					anySymbol = anySymbol || checkLineForSymbol(lines[lineIndex+1], startIndex, endIndex)
				}
				//check right
				if endIndex < len(line)-1 {
					rightChar := rune(line[endIndex])
					anySymbol = anySymbol || isSymbol(rightChar)
				}
				//check left
				if startIndex > 0 {
					leftChar := rune(line[startIndex-1])
					anySymbol = anySymbol || isSymbol(leftChar)
				}
				if anySymbol {
					sum += num
				}
				startIndex = endIndex
			}
		}
	}
	result := strconv.Itoa(sum)
	return &result, nil
}

func (m Day3) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	sum := 0
	for lineIndex, line := range lines {
		for index := 0; index < len(line); index++ {
			if rune(line[index]) == '*' {
				product := 1
				count := 0
				//check above
				if lineIndex > 0 {
					numbersAbove := getNumbers(lines[lineIndex-1], index)
					for _, num := range numbersAbove {
						product *= num
						count++
					}
				}
				//check below
				if lineIndex < len(lines)-1 {
					numbersBelow := getNumbers(lines[lineIndex+1], index)
					for _, num := range numbersBelow {
						product *= num
						count++
					}
				}
				//check right
				if index < len(line)-1 {
					rightChar := rune(line[index+1])
					if unicode.IsDigit(rightChar) {
						endRightIndex := getEndNum(line, index+1)
						rightNum, _ := strconv.Atoi(line[index+1 : endRightIndex])
						product *= rightNum
						count++
					}
				}
				//check left
				if index > 0 {
					leftChar := rune(line[index-1])
					if unicode.IsDigit(leftChar) {
						startLeftIndex := getStartNum(line, index-1)
						leftNum, _ := strconv.Atoi(line[startLeftIndex:index])
						product *= leftNum
						count++
					}
				}
				//Only matters if there are exactly two - don't add the product sum if that doesn't apply
				if count != 2 {
					product = 0
				}
				sum += product
			}

		}
	}
	result := strconv.Itoa(sum)
	return &result, nil
}

func getNumbers(line string, index int) []int {
	var result []int
	centerChar := rune(line[index])
	//can only be one number because if it goes through the center,either one to the left or right HAVE to be the same
	if unicode.IsDigit(centerChar) {
		startIndex := getStartNum(line, index)
		rightIndex := getEndNum(line, index)
		centerNum, _ := strconv.Atoi(line[startIndex:rightIndex])
		result = append(result, centerNum)
	} else {
		if index > 0 {
			leftChar := rune(line[index-1])
			if unicode.IsDigit(leftChar) {
				startIndex := getStartNum(line, index-1)
				leftNum, _ := strconv.Atoi(line[startIndex:index])
				result = append(result, leftNum)
			}
		}
		if index < len(line)-1 {
			rightChar := rune(line[index+1])
			if unicode.IsDigit(rightChar) {
				rightIndex := getEndNum(line, index+1)
				rightNum, _ := strconv.Atoi(line[index+1 : rightIndex])
				result = append(result, rightNum)
			}
		}
	}
	return result
}

// Checks a line for a symbol - start and end correspond to the number above or below
func checkLineForSymbol(line string, start int, end int) bool {
	startIndex := start
	//because diagonals count, expand one to the left if we aren't in a corner
	if startIndex > 0 {
		startIndex--
	}
	endIndex := end
	//same for start, because diagonals count, expand to the right
	if endIndex < len(line)-1 {
		endIndex++
	}
	for i := startIndex; i < endIndex; i++ {
		letter := rune(line[i])
		if isSymbol(letter) {
			return true
		}
	}
	return false
}

func isSymbol(char rune) bool {
	return !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '.'
}

func getEndNum(line string, index int) int {
	for i := index; i < len(line); i++ {
		char := rune(line[i])
		if unicode.IsDigit(char) == false {
			return i
		}
	}
	return len(line)
}

func getStartNum(line string, index int) int {
	for i := index; i >= 0; i-- {
		char := rune(line[i])
		if unicode.IsDigit(char) == false {
			return i + 1
		}
	}
	return 0
}
