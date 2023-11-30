package main

import (
	"fmt"
	"github.com/RichardD012/go-advent-of-code/tasks"
	_ "github.com/RichardD012/go-advent-of-code/tasks/2023"
	"github.com/go-resty/resty/v2"
	"os"
	"path/filepath"
	"time"
)

func createTask(year int, day int) (tasks.DailyTask, error) {
	constructor, ok := tasks.Registry[tasks.TaskKey(year, day)]
	if !ok {
		return nil, fmt.Errorf("no task found for year %d, day %d", year, day)
	}
	return constructor(), nil
}

func main() {
	currentTime := time.Now()
	day := currentTime.Day()
	year := currentTime.Year()
	month := currentTime.Month()

	//Debug override for previous years
	/*if true {
		day = 1
		year = 2021
		month = 12
	}*/

	if month != 12 {
		fmt.Println("Not currently December")
		os.Exit(1)
	}

	data, err := getData(year, day)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		os.Exit(1)
	}

	dailyTask, err := createTask(year, day)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Results for %d, day %d\n", year, day)
	task1Result, err := dailyTask.Task1(data)
	if err != nil {
		fmt.Println("Error processing Task 1: ", err)
		os.Exit(1)
	}
	fmt.Printf("Task 1: %s\n", *task1Result)

	task2Result, err := dailyTask.Task2(data)
	if err != nil {
		fmt.Println("Error processing Task 2: ", err)
		os.Exit(1)
	}
	fmt.Printf("Task 2: %s\n", *task2Result)
}

func getData(year int, day int) (*string, error) {
	filePath := filepath.Join("input", fmt.Sprintf("%d", year), fmt.Sprintf("day-%d-input.txt", day))

	// Read the file if it exists
	if file, err := os.ReadFile(filePath); err == nil {

		result := string(file)
		return &result, nil
	}

	cookie := os.Getenv("AOC_COOKIE")
	if cookie == "" {
		return nil, fmt.Errorf("AOC_COOKIE environment variable is not set")
	}
	client := resty.New()

	client.SetHeader("cookie", cookie)
	resp, err := client.R().Get(fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day))
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v\n", err)
	}
	responseBody := resp.String()

	// Create directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return nil, fmt.Errorf("error creating directory: %v", err)
	}

	if err = os.WriteFile(filePath, []byte(responseBody), 0644); err != nil {
		return nil, fmt.Errorf("error writing file: %v", err)
	}

	return &responseBody, nil
}
