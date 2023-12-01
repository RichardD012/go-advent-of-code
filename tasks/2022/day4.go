package tasks2022

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"strconv"
	"strings"
)

type Day4 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2022, 4), func() tasks.DailyTask { return Day4{} })
}

func (m Day4) Task1(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	sum := 0
	for _, line := range lines {
		set := strings.Split(line, ",")
		group1 := set[0]
		group2 := set[1]
		g1min, g1max := getMinMax(group1)
		g2min, g2max := getMinMax(group2)
		if (g1min <= g2min && g1max >= g2max) || (g2min <= g1min && g2max >= g1max) {
			sum++
		}
	}
	result := strconv.Itoa(sum)
	return &result, nil
}

func (m Day4) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	sum := 0
	for _, line := range lines {
		set := strings.Split(line, ",")
		group1 := set[0]
		group2 := set[1]
		g1min, g1max := getMinMax(group1)
		g2min, g2max := getMinMax(group2)
		if (g1min <= g2min && g1max >= g2max) ||
			(g2min <= g1min && g2max >= g1max) ||
			(g1min <= g2min && g1max >= g2min) ||
			(g2min <= g1min && g2max >= g1min) {
			sum++
		}
	}
	result := strconv.Itoa(sum)
	return &result, nil
}

func getMinMax(group string) (int, int) {
	splitData := strings.Split(group, "-")
	minValue, _ := strconv.Atoi(splitData[0])
	maxValue, _ := strconv.Atoi(splitData[1])
	return minValue, maxValue
}
