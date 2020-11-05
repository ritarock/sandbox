package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	target := strings.Join(os.Args[1:], "")
	output := strings.Split(target, ".")[0] + "_transpose.csv"
	readFile, err := os.Open(target)

	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(readFile)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	writeFile, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(writeFile)
	w.Comma = ','
	w.UseCRLF = true

	result := [][]string{}
	for i := 0; i < len(records[0]); i++ {
		tmp_row := []string{}
		for _, record := range records {
			tmp_row = append(tmp_row, record[i])
		}
		result = append(result, tmp_row)
	}
	fmt.Println(result)

	for _, record := range result {
		if err := w.Write(record); err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
