package main

import (
	"fmt"
	"strconv"
)

func makeTemplate(x int, y string, z float64) string {
	return strconv.Itoa(x) + "時の" + y + "は" + strconv.FormatFloat(z, 'f', -1, 64)
}

func main() {
	fmt.Println(makeTemplate(12, "気温", 22.4))
}
