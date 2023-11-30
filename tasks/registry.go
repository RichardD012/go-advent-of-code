package tasks

import (
	"fmt"
)

type DailyTask interface {
	Task1(*string) (*string, error)
	Task2(*string) (*string, error)
}

var Registry = make(map[string]func() DailyTask)

func RegisterStruct(name string, constructor func() DailyTask) {
	Registry[name] = constructor
}

func TaskKey(year int, day int) string {
	return fmt.Sprintf("%d.Day%d", year, day)
}
