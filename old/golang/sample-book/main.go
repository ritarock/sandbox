package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const BASE_URL = "https://api.openbd.jp/v1"

type OpenBD []struct {
	Onix struct {
		DescriptiveDetail struct {
			Subject []struct {
				SubjectCode string `json:"SubjectCode"`
			} `json:"Subject"`
		} `json:"DescriptiveDetail"`
	} `json:"Onix"`
	Summary struct {
		Isbn      string `json:"isbn"`
		Title     string `json:"title"`
		Publisher string `json:"publisher"`
		Pubdate   string `json:"pubdate"`
		Cover     string `json:"cover"`
		Author    string `json:"author"`
	} `json:"Summary"`
}

type Book struct {
	Isbn        string `json:"isbn"`
	Title       string `json:"title"`
	Publisher   string `json:"publisher"`
	Pubdate     string `json:"pubdate"`
	Cover       string `json:"cover"`
	Author      string `json:"author"`
	SubjectCode string `json:"SubjectCode"`
}

const WORKER_NUM = 10

func main() {
	coverage := getCoverage()
	isbnArr := func(coverage []string) [][]string {
		split := [][]string{}
		sliceSize := len(coverage)
		for i := 0; i < sliceSize; i += 10000 {
			end := i + 10000
			if sliceSize < end {
				end = sliceSize
			}
			split = append(split, coverage[i:end])
		}
		return split
	}(coverage)

	for _, v := range isbnArr {
		getBook(v)
	}
}

func getCoverage() []string {
	var coverage []string
	url := BASE_URL + "/coverage"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&coverage); err != nil {
		fmt.Println(err)
	}
	return coverage
}

func getBook(coverage []string) []Book {
	var openbd OpenBD
	var books []Book
	isbn := strings.Join(coverage, ",")
	urls := BASE_URL + "/get"
	params := url.Values{}
	params.Add("isbn", isbn)
	response, err := http.PostForm(urls, params)

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&openbd); err != nil {
		fmt.Println(err)
	}

	for _, v := range openbd {
		if len(v.Onix.DescriptiveDetail.Subject) == 0 {
			continue
		}
		category := v.Onix.DescriptiveDetail.Subject[0].SubjectCode
		if len(category) != 4 {
			continue
		}
		if strings.Split(category, "")[2]+strings.Split(category, "")[3] != "79" {
			continue
		}

		book := Book{
			Isbn:        v.Summary.Isbn,
			Title:       v.Summary.Title,
			Publisher:   v.Summary.Publisher,
			Pubdate:     v.Summary.Pubdate,
			Cover:       v.Summary.Cover,
			Author:      v.Summary.Author,
			SubjectCode: category,
		}
		fmt.Println(book)
		books = append(books, book)
	}
	return books
}
