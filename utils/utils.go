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
