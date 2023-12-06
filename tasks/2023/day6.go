package tasks2023

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"strconv"
	"strings"
)

type Day6 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2023, 6), func() tasks.DailyTask { return Day6{} })
}

func (m Day6) Task1(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	//lines = m.getTestData()
	timeLine := strings.Split(lines[0], ":")
	distanceLine := strings.Split(lines[1], ":")
	timeParams := strings.Fields(timeLine[1])
	distanceParams := strings.Fields(distanceLine[1])
	var races []Race
	for i := 0; i < len(timeParams); i++ {
		races = append(races, Race{time: getNum(timeParams[i]), distance: getNum(distanceParams[i])})
	}
	raceProduct := 1
	for _, race := range races {
		currentRaceCount := 0
		for i := 0; i <= race.time; i++ {
			time := race.time - i
			speed := i
			distance := time * speed
			if distance > race.distance && time > 0 {
				currentRaceCount++
			}
		}
		raceProduct *= currentRaceCount
	}
	result := strconv.Itoa(raceProduct)
	return &result, nil
}

func (m Day6) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	//lines = m.getTestData()
	timeLine := strings.Split(lines[0], ":")
	distanceLine := strings.Split(lines[1], ":")
	timeParams := strings.Fields(timeLine[1])
	distanceParams := strings.Fields(distanceLine[1])
	var races []Race
	timeString := ""
	distanceString := ""
	for i := 0; i < len(timeParams); i++ {
		timeString += timeParams[i]
		distanceString += distanceParams[i]
	}
	races = append(races, Race{time: getNum(timeString), distance: getNum(distanceString)})
	raceProduct := 1
	for _, race := range races {
		currentRaceCount := 0
		for i := 0; i <= race.time; i++ {
			time := race.time - i
			speed := i
			distance := time * speed
			if distance > race.distance && time > 0 {
				currentRaceCount++
			}
		}
		raceProduct *= currentRaceCount
	}
	result := strconv.Itoa(raceProduct)
	return &result, nil
}

func getNum(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func (m Day6) getTestData() []string {
	return []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
}

type Race struct {
	time     int
	distance int
}
