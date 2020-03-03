package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
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

// getCSVFile search for a file in the system and returns it
func getCSVFile(fileSrc string) *os.File {
	file, error := os.Open(fileSrc)

	if error != nil {
		handleError(error)
	}
	return file
}

func parseCSV(csvFile *os.File) [][]string {
	csvReader := csv.NewReader(csvFile)

	results, error := csvReader.ReadAll()

	if error != nil {
		handleError(error)
	}

	return results
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
	fileFlag := flag.String("Questions' file", "problems.csv", "Sets the CSV file to be used to create the questions")
	// isShuffled := flag.Bool("Shuffle questions", false, "Set wether questions should be shuffled on each game iteration or not")

	file := getCSVFile(*fileFlag)
	questions := parseCSV(file)

	points := askquestions(questions)

	fmt.Print("You got right ", points, " of ", len(questions), "\n")
}
