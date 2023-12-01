package tasks2022

import (
	"fmt"
	"github.com/RichardD012/go-advent-of-code/tasks"
	"github.com/RichardD012/go-advent-of-code/utils"
	"strconv"
	"strings"
)

type Day5 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2022, 5), func() tasks.DailyTask { return Day5{} })
}

func (m Day5) Task1(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	stacks, lastIndex := generateStacks(lines)
	for i := lastIndex + 1; i < len(lines); i++ {
		line := lines[i]
		numMove, from, to := processDirection(line)
		for n := 0; n < numMove; n++ {
			newVal, _ := stacks[from-1].Pop()
			stacks[to-1].Push(newVal)
		}
	}
	result := ""
	for _, entry := range stacks {
		result = fmt.Sprintf("%s%s", result, entry.GetLast())
	}
	return &result, nil
}

func (m Day5) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	stacks, lastIndex := generateStacks(lines)
	for i := lastIndex + 1; i < len(lines); i++ {
		line := lines[i]
		numMove, from, to := processDirection(line)
		var tempStack utils.Stack[string]
		for n := 0; n < numMove; n++ {
			newVal, _ := stacks[from-1].Pop()
			tempStack.Push(newVal)
		}
		for n := 0; n < numMove; n++ {
			newVal, _ := tempStack.Pop()
			stacks[to-1].Push(newVal)
		}

	}
	result := ""
	for _, entry := range stacks {
		result = fmt.Sprintf("%s%s", result, entry.GetLast())
	}
	return &result, nil
}

func processDirection(line string) (int, int, int) {
	line = strings.ReplaceAll(line, "move ", "")
	line = strings.ReplaceAll(line, " from ", ",")
	line = strings.ReplaceAll(line, " to ", ",")
	data := strings.Split(line, ",")
	numMove, _ := strconv.Atoi(data[0])
	from, _ := strconv.Atoi(data[1])
	to, _ := strconv.Atoi(data[2])
	return numMove, from, to
}

func generateStacks(lines []string) ([]utils.Stack[string], int) {
	var stackLines []string
	lastIndex := 0
	for index, line := range lines {
		if line == "" {
			stackLines = lines[0:index]
			lastIndex = index
			break
		}
	}
	bottomLine := strings.Fields(stackLines[len(stackLines)-1])
	numColumns, _ := strconv.Atoi(bottomLine[len(bottomLine)-1])
	stacks := make([]utils.Stack[string], numColumns)
	for i := lastIndex - 2; i >= 0; i-- {
		currentLine := lines[i]
		index := 0
		for n := 1; n < len(currentLine); n += 4 {
			value := string(currentLine[n])
			if value != "" && value != " " {
				stacks[index].Push(value)
			}
			index++
		}
	}
	return stacks, lastIndex
}

func printStacks(stacks []utils.Stack[string]) {
	for _, stack := range stacks {
		stack.Print()
	}
}
