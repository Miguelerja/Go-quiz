package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func handleError(err error) {
	fmt.Print(err)
	os.Exit(1)
}

// setScanner creates a new scanner
func setScanner() (scanner *bufio.Scanner) {
	scanner = bufio.NewScanner(os.Stdin)

	return
}

// getUserInitialConfig gets user inputs to configure the game's time and csv to be used
func getUserInitialConfig() (csvFile string) {
	scanner := setScanner()

	fmt.Print("Enter your custom CSV file. If no file is entered, the default file will be used:\n")
	scanner.Scan()
	csvFile = scanner.Text()

	return
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
		handleError(error)
	}
	return file
}

func parseCSV(csvFile *os.File) (results [][]string) {
	csvReader := csv.NewReader(csvFile)

	for {
		record, error := csvReader.Read()

		if error == io.EOF {
			break
		}
		if error != nil {
			handleError(error)
		}

		results = append(results, record)
	}

	return
}

// askquestions receives a matrix containing questions and answers, loops through the matrix printing the questions
// and gathers hte users input. Once the loop is over, it returns the number of correct answers
func askquestions(questions [][]string) (points int) {
	scanner := setScanner()

	for i := 0; i < len(questions); i++ {
		question := questions[i][0]
		answer := questions[i][1]
		var guess string

		fmt.Print(question + "\n")
		scanner.Scan()
		guess = scanner.Text()

		if guess == answer {
			points = points + 1
		}
	}

	return
}

func main() {
	file := getCSVFile()
	questions := parseCSV(file)

	points := askquestions(questions)

	fmt.Print("You got right ", points, " of ", len(questions), "\n")
}
