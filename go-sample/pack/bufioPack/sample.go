package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := `XXX
YYY
ZZZ`

	r := strings.NewReader(s)

	scanner := bufio.NewScanner(r)

	scanner.Scan()
	fmt.Println(scanner.Text())
	scanner.Scan()
	fmt.Println(scanner.Text())
	scanner.Scan()
	fmt.Println(scanner.Text())
}
