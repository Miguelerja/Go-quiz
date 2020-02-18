package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func getCSVFile() *os.File {
	var csvFile string
	var fileSrc string
	defaultCsvFile := "problems.csv"

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your custom CSV file. If no file is entered, the default file will be used:/n")
	scanner.Scan()
	csvFile = scanner.Text()

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

func main() {
	file := getCSVFile()
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
