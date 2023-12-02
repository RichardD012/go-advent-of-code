package tasks2022

import (
	"fmt"
	"github.com/RichardD012/go-advent-of-code/tasks"
	"strconv"
	"strings"
)

type Day7 struct {
}

func init() {
	tasks.RegisterStruct(tasks.TaskKey(2022, 7), func() tasks.DailyTask { return Day7{} })
}

type Instruction int

const (
	List Instruction = iota
	Change
)

func (m Day7) Task1(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	root := generateDirectoryStructure(lines)
	var suitableDirectories []*Directory
	suitableDirectories = findSmallDirectories(suitableDirectories, root)
	size := 0
	for _, dir := range suitableDirectories {
		size += dir.Size()
	}
	result := strconv.Itoa(size)
	return &result, nil
}

func (m Day7) Task2(data *string) (*string, error) {
	lines := strings.Split(*data, "\n")
	/*lines = []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}*/
	root := generateDirectoryStructure(lines)
	minSize := 30000000 - (70000000 - root.Size())
	dirToDelete := findLargeDirectories(root, nil, minSize)
	result := strconv.Itoa(dirToDelete.Size())
	return &result, nil
}

func findSmallDirectories(directories []*Directory, dir *Directory) []*Directory {
	if dir.Size() <= 100000 {
		directories = append(directories, dir)
	}
	// Recursively iterate through each subdirectory
	for _, subDir := range dir.directories {
		directories = findSmallDirectories(directories, subDir)
	}
	return directories
}

func findLargeDirectories(directory *Directory, currentSmallest *Directory, size int) *Directory {
	if directory.Size() >= size {
		if currentSmallest == nil || currentSmallest.Size() > directory.Size() {
			currentSmallest = directory
		}
	}
	// Recursively iterate through each subdirectory
	for _, subDir := range directory.directories {
		currentSmallest = findLargeDirectories(subDir, currentSmallest, size)
	}
	return currentSmallest
}

func generateDirectoryStructure(lines []string) *Directory {
	root := &Directory{dirName: "/"}
	currentDirectory := root
	for _, line := range lines[1:] {
		if strings.HasPrefix(line, "$ ") {
			//determine which instructions
			instruction, param := getInstruction(line)
			switch instruction {
			case List:
				break
			case Change:
				if param != nil && *param == "/" {
					currentDirectory = root
				} else if param != nil && *param == ".." {
					currentDirectory = currentDirectory.parentDir
				} else if param != nil {
					//find the directory in its children
					set := false
					for _, child := range currentDirectory.directories {
						if child.dirName == *param {
							currentDirectory = child
							set = true
							break
						}
					}
					if set == false {
						fmt.Println("Did not find a child directory")
					}
				} else {
					fmt.Println("Error parsing directory - got a change with no parameters")
				}
				break
			}
		} else {
			if strings.HasPrefix(line, "dir ") {
				newDir := &Directory{dirName: strings.ReplaceAll(line, "dir ", "")}
				newDir.parentDir = currentDirectory
				currentDirectory.directories = append(currentDirectory.directories, newDir)
			} else {
				//it's a file
				splitString := strings.Split(line, " ")
				size, _ := strconv.Atoi(splitString[0])
				subDirFile := File{fileName: splitString[1], size: size}
				currentDirectory.files = append(currentDirectory.files, subDirFile)
			}

		}
	}
	return root
}

func getInstruction(line string) (Instruction, *string) {
	if strings.HasPrefix(line, "$ ls") {
		return List, nil
	}
	//its a change
	cleanedString := strings.ReplaceAll(line, "$ cd ", "")
	return Change, &cleanedString
}

type Directory struct {
	parentDir   *Directory
	dirName     string
	files       []File
	directories []*Directory
}

// Size calculates the total size of the directory,
// including all files and subdirectories
func (d *Directory) Size() int {
	totalSize := 0

	// Sum the size of all files in this directory
	for _, file := range d.files {
		totalSize += file.size
	}

	// Recursively sum the size of all subdirectories
	for _, dir := range d.directories {
		totalSize += dir.Size() // Recursive call
	}

	return totalSize
}

type File struct {
	fileName string
	size     int
}
