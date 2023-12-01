package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// StringToIntSlice converts a string to a slice of ints.
// Assumes the string contains one integer per line.
func StringToIntSlice(data *string) ([]int, error) {

	if data == nil {
		return nil, fmt.Errorf("input string pointer is nil")
	}
	// Split the string into lines
	lines := strings.Split(*data, "\n")

	// Initialize a slice to hold the parsed numbers
	var numbers []int

	// Iterate over the lines and convert each one to an integer
	for _, line := range lines {
		// Skip empty lines
		if line == "" {
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			// Return the error and the numbers parsed so far
			return numbers, fmt.Errorf("error parsing '%s' as integer: %v", line, err)
		}

		// Append the number to the slice
		numbers = append(numbers, num)
	}

	return numbers, nil
}

// Stack type based on a slice of integers
type Stack[T any] []T

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

// Pop removes an element from the top of the stack and returns it
// Returns 0 and false if the stack is empty
func (s *Stack[T]) Pop() (T, bool) {
	if len(*s) == 0 {
		var zero T
		return zero, false
	}
	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]
	return value, true
}

// GetLast just gets the last without popping
func (s *Stack[T]) GetLast() T {
	if len(*s) == 0 {
		var zero T
		return zero
	}
	index := len(*s) - 1
	value := (*s)[index]
	return value
}

func (s *Stack[T]) Print() {
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

type Queue[T any] []T

func (q *Queue[T]) Enqueue(value T) {
	*q = append(*q, value)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(*q) == 0 {
		var zero T
		return zero, false
	}
	value := (*q)[0]
	*q = (*q)[1:]
	return value, true
}
