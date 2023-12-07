package tasks2023

import (
	"github.com/RichardD012/go-advent-of-code/tasks"
	"sort"
	"strconv"
	"strings"
)

type Day7 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2023, 7), func() tasks.DailyTask { return Day7{} })
}

func (m Day7) Task1(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	//lines = m.getTestData()
	var hands []Hand
	for _, line := range lines {
		hand := generateHand(line, false)
		hands = append(hands, hand)
	}
	//sort them
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].hand != hands[j].hand {
			return hands[i].hand > hands[j].hand
		}
		for index, _ := range hands[i].line {
			rankCard1 := getRank(rune(hands[i].line[index]), false)
			rankCard2 := getRank(rune(hands[j].line[index]), false)
			if rankCard1 == rankCard2 {
				continue
			}
			return rankCard1 > rankCard2
		}
		return hands[i].hand > hands[j].hand
	})
	product := 0
	for index, hand := range hands {
		product += (index + 1) * hand.wager
	}
	result := strconv.Itoa(product)
	return &result, nil
}

func (m Day7) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	//lines = m.getTestData()
	var hands []Hand
	for _, line := range lines {
		hand := generateHand(line, true)
		hands = append(hands, hand)
	}
	//sort them
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].hand != hands[j].hand {
			return hands[i].hand > hands[j].hand
		}
		for index, _ := range hands[i].line {
			rankCard1 := getRank(rune(hands[i].line[index]), true)
			rankCard2 := getRank(rune(hands[j].line[index]), true)
			if rankCard1 == rankCard2 {
				continue
			}
			return rankCard1 > rankCard2
		}
		return hands[i].hand > hands[j].hand
	})
	product := 0
	for index, hand := range hands {
		product += (index + 1) * hand.wager
	}
	result := strconv.Itoa(product)
	return &result, nil
}

func getRank(letter rune, partTwo bool) int {
	switch letter {
	case 'A':
		return 0
	case 'K':
		return 1
	case 'Q':
		return 2
	case 'J':
		if partTwo {
			return 20
		}
		return 3
	case 'T':
		return 4
	case '9':
		return 5
	case '8':
		return 6
	case '7':
		return 7
	case '6':
		return 8
	case '5':
		return 9
	case '4':
		return 10
	case '3':
		return 11
	case '2':
		return 12

	}
	return 13
}

func generateHand(line string, day2 bool) Hand {
	split := strings.Split(line, " ")
	wager, _ := strconv.Atoi(split[1])
	letterMap, jokers := generateLetterMap(split[0])
	if day2 && jokers > 0 {
		var replaceChar rune
		var maxCount int
		//if there are any J's go through and replace it with the letter with the highest count
		//this will create a hand with the next best. 2->3, 3->4, 4->5 because all hands go up
		tempLine := split[0]
		for char, count := range letterMap {
			if char == 'J' {
				continue
			}
			if count > maxCount {
				replaceChar = char
				maxCount = count
			}
		}
		//replace the most popular letter with J and regenerate the hand to get the correct rank ie QQQJA turns to QQQQA
		// which is four of a kind
		tempLine = strings.ReplaceAll(tempLine, "J", string(replaceChar))
		letterMap, _ = generateLetterMap(tempLine)
	}
	var handType HandType
	switch len(letterMap) {
	case 5:
		handType = HighCard
		break
	case 4:
		handType = OnePair
		break
	case 3:
		handType = TwoPair
		for _, count := range letterMap {
			if count == 3 {
				handType = ThreeOfKind
			}
		}
		break
	case 2:
		handType = FullHouse
		for _, count := range letterMap {
			if count == 4 {
				handType = FourOfKind
			}
		}
		break
	case 1:
		handType = FiveOfKind
		break
	}
	//return the original hand for tie-break comparison
	return Hand{wager: wager, hand: handType, line: split[0]}

}

func generateLetterMap(line string) (map[rune]int, int) {
	letterMap := make(map[rune]int)
	jokers := 0
	for _, ch := range line {
		letterMap[rune(ch)]++
		if rune(ch) == 'J' {
			jokers++
		}
	}
	return letterMap, jokers
}

func (m Day7) getTestData() []string {
	return []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
}

type Hand struct {
	hand  HandType
	line  string
	wager int
}

type HandType int

const (
	FiveOfKind HandType = iota
	FourOfKind
	FullHouse
	ThreeOfKind
	TwoPair
	OnePair
	HighCard
)
