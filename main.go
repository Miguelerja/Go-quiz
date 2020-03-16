package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
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

// parseCSV reads a csvFile and parses it into a matrix of string arrats containing question and answer
func parseCSV(csvFile *os.File) [][]string {
	csvReader := csv.NewReader(csvFile)
	results, error := csvReader.ReadAll()

	if error != nil {
		handleError(error)
	}

	return results
}

// shuffleArr takes an array and returns it with its components shuffled randomly
func shuffleArr(arr [][]string) [][]string {
	copyArr := append([][]string(nil), arr...)

	for i, j := range rand.Perm(len(arr)) {
		arr[i] = copyArr[j]
	}

	return arr
}

// askquestions receives a matrix containing questions and answers, loops through the matrix printing the questions
// and gathers the users input. Once the loop is over, it returns the number of correct answers
func askquestions(questions [][]string, isShuffled bool) (points int) {
	scanner := setScanner()

	if isShuffled {
		shuffleArr(questions)
	}

	for _, item := range questions {
		question := item[0]
		answer := item[1]
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
	fileFlag := flag.String("questions", "problems.csv", "Sets the CSV file to be used to create the questions")
	isShuffled := flag.Bool("shuffle", false, "Set wether questions should be shuffled on each game iteration or not")

	flag.Parse()

	file := getCSVFile(*fileFlag)
	questions := parseCSV(file)

	points := askquestions(questions, *isShuffled)

	fmt.Print("You got right ", points, " of ", len(questions), "\n")
}
