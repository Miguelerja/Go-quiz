package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// setScanner creates a new scanner, asks the user for input and returns the input value
func setScanner() string {
	var csvFile string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your custom CSV file. If no file is entered, the default file will be used:\n")
	scanner.Scan()
	csvFile = scanner.Text()

	return csvFile
}

// getCSVFile search for a file in the system and returns it
func getCSVFile() *os.File {
	var fileSrc string
	defaultCsvFile := "problems.csv"

	csvFile := setScanner()

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

func parseCSV(file *os.File) {
	csvReader := csv.NewReader(file)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Print(err)
		}

		fmt.Print(record)
	}

}

func main() {
	file := getCSVFile()

	parseCSV(file)
}
