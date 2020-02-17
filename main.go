package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	file, error := os.Open("problems.csv")
	if error != nil {
		fmt.Print(error)
	}

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Print(err)
		}

		fmt.Print(record)
	}
}
