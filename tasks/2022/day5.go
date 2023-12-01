package tasks2022

import (
	"fmt"
	"github.com/RichardD012/go-advent-of-code/tasks"
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
		var tempStack Stack
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
	//fmt.Printf("%s\n", line)
	line = strings.ReplaceAll(line, "move ", "")
	line = strings.ReplaceAll(line, " from ", ",")
	line = strings.ReplaceAll(line, " to ", ",")
	data := strings.Split(line, ",")
	numMove, _ := strconv.Atoi(data[0])
	from, _ := strconv.Atoi(data[1])
	to, _ := strconv.Atoi(data[2])
	return numMove, from, to
}

func generateStacks(lines []string) ([]Stack, int) {
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
	stacks := make([]Stack, numColumns)
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

func printStacks(stacks []Stack) {
	for _, stack := range stacks {
		stack.Print()
	}
}

// Stack type based on a slice of integers
type Stack []string

// Push adds an element to the top of the stack
func (s *Stack) Push(value string) {
	*s = append(*s, value)
}

// Pop removes an element from the top of the stack and returns it
// Returns 0 and false if the stack is empty
func (s *Stack) Pop() (string, bool) {
	if len(*s) == 0 {
		return "", false
	}
	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]
	return value, true
}

// GetLast just gets the last without popping
func (s *Stack) GetLast() string {
	if len(*s) == 0 {
		return ""
	}
	index := len(*s) - 1
	value := (*s)[index]
	return value
}

func (s *Stack) Print() {
	fmt.Printf("Stack: ")
	if len(*s) == 0 {
		fmt.Printf("\n")
		return
	}
	for _, letter := range *s {
		fmt.Printf("%v", letter)
	}
	fmt.Printf("\n")
}
