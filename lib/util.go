package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
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

// UserDataDir returns the redovc data dir of the current user.
func UserDataDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return home + "/.rvc/redovc"
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

func JSONtoCSV(jsonData []byte) string {
	b := &bytes.Buffer{}
	wr := json2csv.NewCSVWriter(b)
	wr.HeaderStyle = json2csv.SlashStyle
	var x []map[string]interface{}

	// unMarshall json
	err := json.Unmarshal(jsonData, &x)
	if err != nil {
		fmt.Printf("Error reading .redovc.todos.json file: %v\n", err)
	}

	// convert json to CSV
	csv, err := json2csv.JSON2CSV(x)
	if err != nil {
		fmt.Printf("Error converting json to csv: %v\n", err)
	}

	// CSV bytes convert & writing...
	err = wr.WriteCSV(csv)
	if err != nil {
		fmt.Printf("Error displaying csv: %v\n", err)
	}
	wr.Flush()
	got := b.String()

	//Following line prints CSV
	return got
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

// JSON2HtmlTable convert json string to html table string
func JSON2HtmlTable(jsonStr string, customTitles []string, rowSpanTitles []string) (bool, string) {
	htmlTable := ""
	jsonArray := []map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonStr), &jsonArray)
	if err != nil || len(jsonArray) == 0 {
		fmt.Println("invalid json string")
		return false, htmlTable
	}

	titles := customTitles
	if nil == customTitles || len(customTitles) == 0 {
		titles = getKeys(jsonArray[0])
	}

	if nil != rowSpanTitles && len(rowSpanTitles) != 0 {
		for tid, title := range rowSpanTitles {
			swapped := true
			for swapped {
				swapped = false
				for i := 0; i < len(jsonArray)-1; i++ {
					va, oka := jsonArray[i][title].(string)
					vb, okb := jsonArray[i+1][title].(string)
					if !oka || !okb {
						swapped = false
						break
					}
					if strings.Compare(va, vb) > 0 {
						if tid != 0 {
							va, _ := jsonArray[i][rowSpanTitles[tid-1]].(string)
							vb, _ := jsonArray[i+1][rowSpanTitles[tid-1]].(string)
							if va != vb {
								continue
							}
						}
						tmp := jsonArray[i]
						jsonArray[i] = jsonArray[i+1]
						jsonArray[i+1] = tmp
						swapped = true
					}
				}
			}
		}
	}
	// convert table headers
	if len(titles) == 0 {
		fmt.Println("json is not supported")
	}
	tmp := []string{}
	for _, title := range titles {
		tmp = append(tmp, fmt.Sprintf("<th style='font-weight:bold;background-color:lightgrey'>%s</th>", strings.ToUpper(title)))
	}
	thCon := strings.Join(tmp, "")

	// convert table cells
	segs := map[string][]int{}
	initSeg := []int{0, len(jsonArray)}
	for i, key := range rowSpanTitles {
		seg := initSeg
		for j := 1; j < len(jsonArray); j++ {
			if jsonArray[j][key] != jsonArray[j-1][key] {
				inSlice := false
				for _, k := range seg {
					if k == j {
						inSlice = true
					}
				}
				if !inSlice {
					seg = append(seg, j)
				}
			}
		}
		sort.Ints(seg)
		segs[rowSpanTitles[i]] = seg
		if i < len(rowSpanTitles)-1 {
			segs[rowSpanTitles[i+1]] = segs[key]
			initSeg = segs[key]
		}
	}
	rows := []string{}
	for i, jsonObj := range jsonArray {
		tmp = []string{}
		for _, key := range titles {
			seg := segs[key]
			if len(seg) != 0 {
				if i == 0 {
					cell := fmt.Sprintf(`<td rowspan="%d">%v</td>`, seg[1], jsonObj[key])
					tmp = append(tmp, cell)
				} else {
					for n, j := range seg {
						if j == i {
							rowspan := 1
							if n < len(seg)-1 {
								rowspan = seg[n+1] - seg[n]
							}
							cell := fmt.Sprintf(`<td rowspan="%d">%v</td>`, rowspan, jsonObj[key])
							tmp = append(tmp, cell)
						}
					}
				}
			} else {
				cell := fmt.Sprintf("<td>%v</td>", jsonObj[key])
				tmp = append(tmp, cell)
			}
		}
		tdCon := strings.Join(tmp, "")
		row := fmt.Sprintf("<tr>%s</tr>", tdCon)

		rows = append(rows, row)
	}
	trCon := strings.Join(rows, "")

	htmlTable = fmt.Sprintf(`<table border="0" cellpadding="1" cellspacing="1">%s%s</table>`, fmt.Sprintf("<thead>%s</thead>", thCon), fmt.Sprintf("<tbody>%s</tbody>", trCon))
	return true, htmlTable
}

func getKeys(jsonObj map[string]interface{}) []string {
	keys := make([]string, 0, len(jsonObj))
	for k := range jsonObj {
		keys = append(keys, k)
	}
	return keys
}
