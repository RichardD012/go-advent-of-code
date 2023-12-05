package tasks2023

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Day5 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2023, 5), func() tasks.DailyTask { return Day5{} })
}

func (m Day5) Task1(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	seedStrings := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	seedMap := generateSeedMap(lines)
	minValue := math.MaxInt
	for _, seedString := range seedStrings {
		seed, _ := strconv.Atoi(seedString)
		seed = calculateSeed(seed, seedMap)
		minValue = min(minValue, seed)
	}
	result := strconv.Itoa(minValue)
	return &result, nil
}

func (m Day5) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	//lines = m.getTestData()
	var seedRanges []SeedRange
	seedStringRange := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	for i := 0; i < len(seedStringRange); i = i + 2 {
		minRange, _ := strconv.Atoi(seedStringRange[i])
		rangeCount, _ := strconv.Atoi(seedStringRange[i+1])
		newRange := SeedRange{minRange: minRange, maxRange: minRange + rangeCount - 1}
		seedRanges = append(seedRanges, newRange)
	}
	seedMap := generateSeedMap(lines)
	minValue := math.MaxInt
	//The goal is to take an input range, determine any overlap on destination ranges, and then pass that on
	//iterate through each of the "source" seed ranges
	sort.Slice(seedRanges, func(i, j int) bool {
		return seedRanges[i].minRange < seedRanges[j].minRange
	})
	for _, seedRange := range seedRanges {
		sourceRanges := []SeedRange{seedRange}
		var destRanges []SeedRange

		for _, mapping := range seedMap {
			for _, sourceRange := range sourceRanges {
				found := false
				for _, destRange := range mapping.ranges {
					//if the source is lower than the sorted mapping, it's not in there -> just pass it forward
					if sourceRange.maxRange < destRange.source {
						destRanges = append(destRanges, sourceRange)
						found = true
						break
						//the source range starts before the mapped range and ends somewhere inside it
					} else if sourceRange.minRange < destRange.source && sourceRange.maxRange >= destRange.source && sourceRange.maxRange <= (destRange.source+destRange.count) {
						//create new pass forward range that contains ones that don't exist from the source PLUS the ones that do from the dest range
						destRanges = append(destRanges, SeedRange{minRange: sourceRange.minRange, maxRange: destRange.source - 1})
						destRanges = append(destRanges, SeedRange{minRange: destRange.dest, maxRange: destRange.dest + (sourceRange.maxRange - destRange.source)})
						found = true
						break
						//entire source range lives inside of the destination range
					} else if sourceRange.minRange >= destRange.source && sourceRange.maxRange < (destRange.source+destRange.count) {
						destRanges = append(destRanges, SeedRange{minRange: destRange.dest + (sourceRange.minRange - destRange.source),
							maxRange: destRange.dest + (sourceRange.minRange - destRange.source) + (sourceRange.maxRange - sourceRange.minRange)})
						found = true
						break
						//source range goes past the current map
					} else if sourceRange.minRange >= destRange.source && sourceRange.minRange < (destRange.source+destRange.count-1) && sourceRange.maxRange >= (destRange.source+destRange.count) {
						destRanges = append(destRanges, SeedRange{minRange: destRange.dest + (sourceRange.minRange - destRange.source),
							maxRange: destRange.dest + destRange.count - 1})
						sourceRange.minRange = destRange.source + destRange.count
						//found is not equal to false because sourceRange NOW is reset to after this and there could be entries after this - if there aren't, then the remainder
						//of this list will get added below
					}

				}
				//Past the end of the list
				if found == false {
					destRanges = append(destRanges, sourceRange)
				}
			}
			//all of our destination ranges can be passed to the next "map"
			sourceRanges = destRanges
			sort.Slice(sourceRanges, func(i, j int) bool {
				return sourceRanges[i].minRange < sourceRanges[j].minRange
			})
			destRanges = make([]SeedRange, 0)

		}
		//sourceRanges is now a list of "locations" - find the min
		for _, entry := range sourceRanges {
			minValue = min(minValue, entry.minRange)
		}
	}

	result := strconv.Itoa(minValue)
	return &result, nil
}

// Task2Alt Brute force method
func (m Day5) Task2Alt(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")

	var seedRanges []SeedRange
	seedStringRange := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	for i := 0; i < len(seedStringRange); i = i + 2 {
		minRange, _ := strconv.Atoi(seedStringRange[i])
		rangeCount, _ := strconv.Atoi(seedStringRange[i+1])
		newRange := SeedRange{minRange: minRange, maxRange: minRange + rangeCount - 1}
		seedRanges = append(seedRanges, newRange)
	}
	seedMap := generateSeedMap(lines)
	var wg sync.WaitGroup
	minValueChannel := make(chan int)
	for _, seedRange := range seedRanges {
		wg.Add(1)
		go func(sr SeedRange) {
			defer wg.Done()
			localMin := math.MaxInt
			for i := sr.minRange; i <= sr.maxRange; i++ {
				localMin = min(localMin, calculateSeed(i, seedMap))
			}
			minValueChannel <- localMin
		}(seedRange)
	}
	// Close channel when all goroutines are done
	go func() {
		wg.Wait()
		close(minValueChannel)
	}()
	minValue := math.MaxInt
	for val := range minValueChannel {
		minValue = min(minValue, val)
	}

	result := strconv.Itoa(minValue)
	return &result, nil
}

func generateSeedMap(lines []string) []Mapping {
	var seedMap []Mapping
	currentMap := Mapping{ranges: []ConvMap{}}
	mapIndex := 0
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		if "" == strings.TrimSpace(line) {
			seedMap = append(seedMap, currentMap)
			mapIndex++
			currentMap = Mapping{ranges: []ConvMap{}}
		} else if strings.Contains(line, " map:") {
			currentMap.title = line
			continue
		} else {
			nums := strings.Split(line, " ")
			dest, _ := strconv.Atoi(nums[0])
			source, _ := strconv.Atoi(nums[1])
			count, _ := strconv.Atoi(nums[2])
			currentMap.ranges = append(currentMap.ranges, ConvMap{source: source, dest: dest, count: count})
		}
	}
	seedMap = append(seedMap, currentMap)
	for _, entry := range seedMap {
		sort.Slice(entry.ranges, func(i, j int) bool {
			return entry.ranges[i].source < entry.ranges[j].source
		})
	}
	return seedMap
}

func (m Day5) getTestData() []string {
	return []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}
}

func calculateSeed(seed int, seedMap []Mapping) int {
	for _, currMap := range seedMap {
		seed = findInMap(currMap, seed)
	}
	return seed
}

func findInMap(currMap Mapping, seed int) int {
	for _, curRange := range currMap.ranges {
		if seed >= curRange.source && seed < (curRange.source+curRange.count) {
			offset := seed - curRange.source
			return curRange.dest + offset
		}
	}
	return seed
}

type Mapping struct {
	title  string
	ranges []ConvMap
}

type ConvMap struct {
	source int
	dest   int
	count  int
}

type SeedRange struct {
	minRange int
	maxRange int
}
