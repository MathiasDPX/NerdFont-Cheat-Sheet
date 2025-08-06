package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
	_ "embed"
)

//go:embed nerdfont.csv
var charsCSV string

func readCsvFile() [][]string {
    reader := csv.NewReader(strings.NewReader(charsCSV))
    records, err := reader.ReadAll()
    if err != nil {
        log.Fatalf("Unable to parse embedded chars.csv: %v", err)
    }
    return records
}

func LeftPad(s string, n int) string {
	if len(s) > n {
		return s
	}

	return strings.Repeat(string(' '), n-len(s)) + s
}

func main() {
	records := readCsvFile()
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Expecting one or more argument")
		return
	}

	catches := [][]string{}
	longest := 0
	for _, record := range records {
		if strings.Contains(record[0], args[0]) {
			name := string(record[0][3:])

			if utf8.RuneCountInString(name) > longest {
				longest = utf8.RuneCountInString(name)
			}

			record[0] = record[0][3:]
			catches = append(catches, record)
		}
	}

	for _, catch := range catches {
		fmt.Printf("%s | %s | %s \n",
			LeftPad(catch[0], longest),
			LeftPad(catch[1], 6),
			catch[2],
		)
	}
}
