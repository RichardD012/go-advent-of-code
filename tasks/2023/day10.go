package tasks2023

import (
	"fmt"
	"github.com/RichardD012/go-advent-of-code/tasks"
	"strconv"
	"strings"
)

type Day10 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2023, 10), func() tasks.DailyTask { return Day10{} })
}

func (m Day10) Task1(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	lines = m.getTestData()
	currentNode := generateLoop(lines)
	startingNode := currentNode
	//loop through and find the longest cycle
	count := 0
	for {
		fmt.Printf("Count: %d\n", count)
		count++
		currentNode = currentNode.next
		if currentNode == nil || currentNode == startingNode {
			break
		}
	}
	result := strconv.Itoa(len(lines))
	return &result, nil
}

func generateLoop(lines []string) *Node {
	returnNode := Node{}
	startingCol := 0
	startingRow := 0
	found := false
	for rowIndex, row := range lines {
		for colIndex, column := range row {
			if column == 'S' {
				startingCol = colIndex
				startingRow = rowIndex
				returnNode.col = colIndex
				returnNode.row = rowIndex
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	currentNode := returnNode
	fmt.Printf("Starting: %d,%d\n", startingCol, startingRow)
	currentCol := startingCol
	currentRow := startingRow
	for {
		currentChar := rune(lines[currentRow][currentCol])
		nextNode := Node{char: currentChar, row: currentRow, col: currentCol, prev: &currentNode}
		switch currentChar {
		case '|':

			break
		case '-':

			break
		case 'L':
			//north and east
			break
		case 'J':
			//north and wes

			break
		case '7':
			//south and west
			//did I come from the south or did I come from the west
			break
		case 'F':
			//south and east

			break
		}
		currentNode.next = &nextNode
		currentNode = nextNode
		if currentCol == startingCol && currentRow == startingRow {
			currentNode.next = &returnNode
			returnNode.prev = &currentNode
			break
		}
	}
	return &returnNode
}

func getStartingPosition(x, y int, input []string) (int, int) {

	if y != (len(input[x])-1) && (input[x][y+1] == '-' || input[x][y+1] == 'J' || input[x][y+1] == '7') {
		y++
		return x, y
	}

	if x != (len(input)-1) && (input[x+1][y] == '|' || input[x+1][y] == 'J' || input[x+1][y] == 'L') {
		x++
		return x, y
	}

	if y != 0 && (input[x][y-1] == 'F' || input[x][y-1] == '-' || input[x][y-1] == 'L') {
		y--
		return x, y
	}

	if x != 0 && (input[x-1][y] == '|' || input[x-1][y] == 'L' || input[x-1][y] == '7') {
		x--
		return x, y
	}
	return x, y
}

func (m Day10) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	result := strconv.Itoa(len(lines))
	return &result, nil
}

func (m Day10) getTestData() []string {
	return []string{
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		".....",
	}
}

type Node struct {
	char rune
	row  int
	col  int
	next *Node
	prev *Node
}
