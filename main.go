package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func getCSVFile() *os.File {
	inputReader := bufio.NewReader(os.Stdin)
	text, _ := inputReader.ReadString('\n')
	var fileSrc string
	if len(os.Args) > 1 {
		fileSrc = os.Args[1]
	} else {
		fileSrc = "problems.csv"
	}

	file, error := os.Open(fileSrc)
	fmt.Print(error)
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
