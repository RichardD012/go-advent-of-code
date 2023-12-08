package tasks2023

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"strconv"
	"strings"
)

type Day8 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2023, 8), func() tasks.DailyTask { return Day8{} })
}

func (m Day8) Task1(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	lineInstructions := lines[0]
	lineMap, _ := generateMap(lines[2:])
	zFound := false
	count := 0
	keyLoop := make(map[string]bool)
	currentKey := "AAA"
	for {
		for _, instruction := range lineInstructions {
			switch rune(instruction) {
			case 'L', 'l':
				currentKey = lineMap[currentKey].left
				break
			case 'R', 'r':
				currentKey = lineMap[currentKey].right
				break
			}
			count++
			if currentKey == "ZZZ" {
				zFound = true
				break
			}
		}
		_, exists := keyLoop[currentKey]
		if exists {
			zFound = true
		}
		if zFound {
			break
		}
	}
	result := strconv.Itoa(count)
	return &result, nil
}

func (m Day8) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	//lines = m.getTestData()
	lineInstructions := lines[0]
	lineMap, currentKeys := generateMap(lines[2:])
	zFound := false
	lcmResult := 1
	for _, currentKey := range currentKeys {
		currentKeyCount := 0
		zFound = false
		for {

			for _, instruction := range lineInstructions {
				switch rune(instruction) {
				case 'L', 'l':
					currentKey = lineMap[currentKey].left
					break
				case 'R', 'r':
					currentKey = lineMap[currentKey].right
					break
				}
				currentKeyCount++
				if currentKey[2] == 'Z' {
					zFound = true
					break
				}
			}
			if zFound {
				break
			}
		}
		lcmResult = lcm(lcmResult, currentKeyCount)
	}
	result := strconv.Itoa(lcmResult)
	return &result, nil
}

// GCD via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find LCM via GCD - from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func lcm(a, b int) int {
	result := a * b / gcd(a, b)
	return result
}

func generateMap(lines []string) (map[string]Next, []string) {
	returnMap := make(map[string]Next)

	var firstKeys []string
	for _, line := range lines {
		splitLine := strings.Split(line, " = ")
		key := splitLine[0]
		coords := strings.Split(splitLine[1], ", ")
		left := coords[0][1:]
		right := coords[1][0 : len(coords[1])-1]
		returnMap[key] = Next{left: left, right: right}
		if rune(key[2]) == 'A' {
			firstKeys = append(firstKeys, key)
		}
	}
	return returnMap, firstKeys
}

type Next struct {
	left  string
	right string
}

func (m Day8) getTestData() []string {
	return []string{
		/*"RL",
		"",
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",*/
		/*"LLR",
		"",
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",*/
		"LR",
		"",
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	}
}
