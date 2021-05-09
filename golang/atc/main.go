package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

type ReadLine interface {
	getI() int
	getS() string
	getF() float64
	getLine() string
	getArrI() []int
}
type TReadLine struct {
}

func solve() {
	// rl := TReadLine{}
}

func sToI(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func sToF(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

func max(arr interface{}) interface{} {
	var v interface{}
	switch arr.(type) {
	case []int:
		newArr := sortReversArr(arr)
		v = newArr.([]int)[0]
	case []float64:
		newArr := sortReversArr(arr)
		v = newArr.([]float64)[0]
	}
	return v
}

func min(arr interface{}) interface{} {
	var v interface{}
	switch arr.(type) {
	case []int:
		newArr := sortArr(arr)
		v = newArr.([]int)[0]
	case []float64:
		newArr := sortArr(arr)
		v = newArr.([]float64)[0]
	}
	return v
}

func (t TReadLine) getI() int {
	var a int
	fmt.Scan(&a)
	return a
}
func (t TReadLine) getS() string {
	var a string
	fmt.Scan(&a)
	return a
}
func (t TReadLine) getF() float64 {
	var a float64
	fmt.Scan(&a)
	return a
}

func (t TReadLine) getLine() string {
	sc.Scan()
	return sc.Text()
}

func (t TReadLine) getArrI() []int {
	var arr []int
	sc.Scan()
	for _, v := range strings.Split(sc.Text(), " ") {
		i, _ := strconv.Atoi(v)
		arr = append(arr, i)
	}
	return arr
}

func (t TReadLine) getArrS() []string {
	sc.Scan()
	return strings.Split(sc.Text(), " ")
}

func (t TReadLine) getArrF() []float64 {
	var arr []float64
	sc.Scan()
	for _, v := range strings.Split(sc.Text(), " ") {
		i, _ := strconv.ParseFloat(v, 64)
		arr = append(arr, i)
	}
	return arr
}

func sortArr(arr interface{}) interface{} {
	switch arr.(type) {
	case []int:
		sort.Sort(sort.IntSlice(arr.([]int)))
	case []float64:
		sort.Sort(sort.Float64Slice(arr.([]float64)))
	}
	return arr
}

func sortReversArr(arr interface{}) interface{} {
	switch arr.(type) {
	case []int:
		sort.Sort(sort.Reverse(sort.IntSlice(arr.([]int))))
	case []float64:
		sort.Sort(sort.Reverse(sort.Float64Slice(arr.([]float64))))
	}
	return arr
}

func uniqueArr(arr interface{}) interface{} {
	var newArr interface{}
	switch arr.(type) {
	case []int:
		m := make(map[int]struct{})
		newArr = make([]int, 0)

		for _, v := range arr.([]int) {
			if _, ok := m[v]; !ok {
				m[v] = struct{}{}
				newArr = append(newArr.([]int), v)
			}
		}
	case []float64:
		m := make(map[float64]struct{})
		newArr = make([]float64, 0)

		for _, v := range arr.([]float64) {
			if _, ok := m[v]; !ok {
				m[v] = struct{}{}
				newArr = append(newArr.([]float64), v)
			}
		}
	}
	return newArr
}

func main() {
	// go run main.go < in.txt
	solve()
}
