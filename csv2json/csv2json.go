package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Data struct {
	Crim    string `json:"crim"`
	Zn      string `json:"zn"`
	Indus   string `json:"indus"`
	Chas    string `json:"chas"`
	Nox     string `json:"nox"`
	Rm      string `json:"rm"`
	Age     string `json:"age"`
	Dis     string `json:"dis"`
	Rad     string `json:"rad"`
	Tax     string `json:"tax"`
	Ptratio string `json:"ptratio"`
	B       string `json:"b"`
	Lstat   string `json:"lstat"`
	Medv    string `json:"medv"`
}

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
	var data = Data{}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else {
			failOnError(err)
		}

		data.Crim = record[0]
		data.Zn = record[1]
		data.Indus = record[2]
		data.Chas = record[3]
		data.Nox = record[4]
		data.Rm = record[5]
		data.Age = record[6]
		data.Dis = record[7]
		data.Rad = record[8]
		data.Tax = record[9]
		data.Ptratio = record[10]
		data.B = record[11]
		data.Lstat = record[12]
		data.Medv = record[13]
		// log.Printf("%#v", record)

		outputJson, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(outputJson))

	}
}
