package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ReadLine interface {
	getInt() int
	getFloat() float64
	getString() string
	getLine() string
	getIntArray() []int
	getFloatArray() []float64
	getStringArray() []string
}

type TReadLine struct {
}

var sc = bufio.NewScanner(os.Stdin)

func solve() {
	// rl := TReadLine{}
}

func toInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func toFloat(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

func (t TReadLine) getInt() int {
	var a int
	fmt.Scan(&a)
	return a
}

func (t TReadLine) getFloat() float64 {
	var a float64
	fmt.Scan(&a)
	return a
}

func (t TReadLine) getString() string {
	var a string
	fmt.Scan(&a)
	return a
}

func (t TReadLine) getLine() string {
	sc.Scan()
	return sc.Text()
}

func (t TReadLine) getIntArray() []int {
	sc.Scan()
	var arr []int
	for _, v := range strings.Split(sc.Text(), " ") {
		i, _ := strconv.Atoi(v)
		arr = append(arr, i)
	}
	return arr
}

func (t TReadLine) getFloatArray() []float64 {
	sc.Scan()
	var arr []float64
	for _, v := range strings.Split(sc.Text(), " ") {
		i, _ := strconv.ParseFloat(v, 64)
		arr = append(arr, i)
	}
	return arr
}

func (t TReadLine) getStringArray() []string {
	sc.Scan()
	return strings.Split(sc.Text(), " ")
}

func SortArray(array interface{}) interface{} {
	switch array.(type) {
	case []float64:
		sort.Sort(sort.Float64Slice(array.([]float64)))
	case []int:
		sort.Sort(sort.IntSlice(array.([]int)))
	}
	return array
}

func ReverseSortArray(array interface{}) interface{} {
	switch array.(type) {
	case []int:
		sort.Sort(sort.Reverse(sort.IntSlice(array.([]int))))
	case []float64:
		sort.Sort(sort.Reverse(sort.Float64Slice(array.([]float64))))
	}
	return array
}

func UniqueArray(array interface{}) interface{} {
	var newList interface{}
	switch array.(type) {
	case []int:
		m := make(map[int]struct{})
		newList = make([]int, 0)

		for _, v := range array.([]int) {
			if _, ok := m[v]; !ok {
				m[v] = struct{}{}
				newList = append(newList.([]int), v)
			}
		}

	case []float64:
		m := make(map[float64]struct{})
		newList = make([]float64, 0)

		for _, v := range array.([]float64) {
			if _, ok := m[v]; !ok {
				m[v] = struct{}{}
				newList = append(newList.([]float64), v)
			}
		}
	}

	return newList
}

func UnionArray(array ...interface{}) interface{} {
	var unionArray interface{}

	switch array[0].(type) {
	case []int:
		m := make(map[int]struct{})
		for _, arr := range array {
			for _, v := range arr.([]int) {
				m[v] = struct{}{}
			}
		}
		tmp := []int{}
		for k := range m {
			tmp = append(tmp, k)
		}
		unionArray = tmp

	case []float64:
		m := make(map[float64]struct{})
		for _, arr := range array {
			for _, v := range arr.([]float64) {
				m[v] = struct{}{}
			}
		}

		tmp := []float64{}
		for k := range m {
			tmp = append(tmp, k)
		}
		unionArray = tmp
	}

	return SortArray(unionArray)
}

func IntersectArray(array ...interface{}) interface{} {
	var intersectArray interface{}

	switch array[0].(type) {
	case []int:
		switch length := len(array); length {
		case 1:
			intersectArray = array[0].([]int)
		case 2:
			tmp := []int{}
			m := make(map[int]struct{})
			for _, v := range array[0].([]int) {
				m[v] = struct{}{}
			}
			for _, v := range array[1].([]int) {
				if _, ok := m[v]; !ok {
					continue
				}
				tmp = append(tmp, v)
			}
			intersectArray = tmp
		default:
			tmp := []int{}
			firstArr := IntersectArray(array[0], array[1])

			for i := 0; i < len(array)-2; i++ {
				func(arr1, arr2 []int) {
					m := make(map[int]struct{})
					for _, v := range arr1 {
						m[v] = struct{}{}
					}
					for _, v := range arr2 {
						if _, ok := m[v]; !ok {
							continue
						}
						tmp = append(tmp, v)
					}
				}(firstArr.([]int), array[i+2].([]int))
			}
			intersectArray = UniqueArray(tmp)
		}

	case []float64:
		switch length := len(array); length {
		case 1:
			intersectArray = array[0].([]float64)
		case 2:
			tmp := []float64{}
			m := make(map[float64]struct{})
			for _, v := range array[0].([]float64) {
				m[v] = struct{}{}
			}
			for _, v := range array[1].([]float64) {
				if _, ok := m[v]; !ok {
					continue
				}
				tmp = append(tmp, v)
			}
			intersectArray = tmp
		default:
			tmp := []float64{}
			firstArr := IntersectArray(array[0], array[1])

			for i := 0; i < len(array)-2; i++ {
				func(arr1, arr2 []float64) {
					m := make(map[float64]struct{})
					for _, v := range arr1 {
						m[v] = struct{}{}
					}
					for _, v := range arr2 {
						if _, ok := m[v]; !ok {
							continue
						}
						tmp = append(tmp, v)
					}
				}(firstArr.([]float64), array[i+2].([]float64))
			}

			intersectArray = UniqueArray(tmp)
		}
	}

	return SortArray(intersectArray)
}

func DifferenceArray(array1, array2 interface{}) interface{} {
	var differenceArray interface{}

	switch array1.(type) {
	case []int:
		m := make(map[int]struct{})
		for _, v := range array2.([]int) {
			m[v] = struct{}{}
		}

		tmp := []int{}
		for _, v := range array1.([]int) {
			if _, ok := m[v]; ok {
				continue
			}
			tmp = append(tmp, v)
		}
		differenceArray = tmp
	case []float64:
		m := make(map[float64]struct{})
		for _, v := range array2.([]float64) {
			m[v] = struct{}{}
		}

		tmp := []float64{}
		for _, v := range array1.([]float64) {
			if _, ok := m[v]; ok {
				continue
			}
			tmp = append(tmp, v)
		}
		differenceArray = tmp

	}

	return differenceArray
}

func main() {
	// go run main.go < in.txt
	solve()
}
