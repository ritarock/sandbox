package main

import "os"

func main() {
	src := "./test.txt"
	dest := "./test.bak"
	_ = os.Link(src, dest)
}
