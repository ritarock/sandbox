package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func main() {
	file, err := os.Open("./boston.csv")
	failOnError(err)
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else {
			failOnError(err)
		}

		log.Printf("%#v", record)
	}
}
