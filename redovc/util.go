package redovc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/twinj/uuid"
	"github.com/yukithm/json2csv"
)

// AddIfNotThere is appending an item to an array if the item is not already present.
func AddIfNotThere(arr []string, items []string) []string {
	for _, item := range items {
		there := false
		for _, arrItem := range arr {
			if item == arrItem {
				there = true
			}
		}
		if !there {
			arr = append(arr, item)
		}
	}
	return arr
}

// AddTodoIfNotThere is appending an todo item to an todo array if the item is not already present.
func AddTodoIfNotThere(arr []*Todo, item *Todo) []*Todo {
	there := false
	for _, arrItem := range arr {
		if item.ID == arrItem.ID {
			there = true
		}
	}
	if !there {
		arr = append(arr, item)
	}
	return arr
}

// UserHomeDir returns the home dir of the current user.
func UserHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return home
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func newUUID() string {
	return (uuid.NewV4()).String()
}

func adjustedDateWithoutTime(t time.Time) time.Time {
	year, month, day := t.Date()

	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func timestamp(t time.Time) time.Time {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()

	return time.Date(year, month, day, hour, min, sec, 0, t.Location())
}

func pluralize(count int, singular, plural string) string {
	if count > 1 {
		return plural
	}
	return singular
}

func isToday(t time.Time) bool {
	nowYear, nowMonth, nowDay := time.Now().Date()
	timeYear, timeMonth, timeDay := t.Date()
	return nowYear == timeYear &&
		nowMonth == timeMonth &&
		nowDay == timeDay
}

func isTomorrow(t time.Time) bool {
	nowYear, nowMonth, nowDay := time.Now().AddDate(0, 0, 1).Date()
	timeYear, timeMonth, timeDay := t.Date()
	return nowYear == timeYear &&
		nowMonth == timeMonth &&
		nowDay == timeDay
}

func isPastDue(t time.Time) bool {
	return time.Now().After(t)
}

func JSONtoCSV(jsonData []byte) {
	b := &bytes.Buffer{}
	wr := json2csv.NewCSVWriter(b)
	var x []map[string]interface{}

	// unMarshall json
	err := json.Unmarshal(jsonData, &x)
	if err != nil {
		fmt.Printf("Error reading .todos.json file: %v\n", err)
	}

	// convert json to CSV
	csv, err := json2csv.JSON2CSV(x)
	if err != nil {
		fmt.Printf("Error reading .todos.json file: %v\n", err)
	}

	// CSV bytes convert & writing...
	err = wr.WriteCSV(csv)
	if err != nil {
		fmt.Printf("Error reading .todos.json file: %v\n", err)
	}
	wr.Flush()
	got := b.String()

	//Following line prints CSV
	fmt.Println(got)
}

func RemoveFromStringArray(array []string, indexToDelete int) []string {
	// check if the index is within array bounds
	if indexToDelete < 0 || indexToDelete >= len(array) {
		return array
	} else {
		// delete an element from the array
		newLength := 0
		for index := range array {
			if indexToDelete != index {
				array[newLength] = array[index]
				newLength++
			}
		}
		return array[:newLength]
	}
}
