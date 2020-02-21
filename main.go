package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// setScanner creates a new scanner
func setScanner() *bufio.Scanner {
	scanner := bufio.NewScanner(os.Stdin)

	return scanner
}

// getUserInitialConfig gets user inputs to configure the game's time and csv to be used
func getUserInitialConfig() string {
	var csvFile string
	scanner := setScanner()

	fmt.Print("Enter your custom CSV file. If no file is entered, the default file will be used:\n")
	scanner.Scan()
	csvFile = scanner.Text()

	return csvFile
}

// getCSVFile search for a file in the system and returns it
func getCSVFile() *os.File {
	var fileSrc string
	defaultCsvFile := "problems.csv"

	csvFile := getUserInitialConfig()

	if len(csvFile) > 0 {
		fileSrc = csvFile
	} else {
		fileSrc = defaultCsvFile
	}

	file, error := os.Open(fileSrc)

	if error != nil {
		fmt.Print(error)
	}
	return file
}

func parseCSV(csvFile *os.File) [][]string {
	csvReader := csv.NewReader(csvFile)
	var results [][]string

	for {
		record, error := csvReader.Read()

		if error == io.EOF {
			break
		}
		if error != nil {
			fmt.Print(error)
			break
		}

		results = append(results, record)
	}

	return results
}

func main() {
	// points := 0
	file := getCSVFile()
	questions := parseCSV(file)

	for i := 0; i < len(questions); i++ {
		question := questions[i][0]
		answer := questions[i][1]
		fmt.Print(question)
		fmt.Print(answer)
	}
}
