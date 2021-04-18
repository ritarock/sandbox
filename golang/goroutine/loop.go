package main

import (
	"fmt"
	"sync"
)

type Item struct {
	Id   int
	Name string
}

func main() {
	list := make([]Item, 10)
	execLoop(list)
}

func execLoop(list []Item) {
	var wg sync.WaitGroup
	for _, item := range list {
		wg.Add(1)
		go func(item2 Item) {
			defer wg.Done()
			do_something(item2)
			wg.Done()
		}(item)
	}
	wg.Wait()
}

func do_something(item Item) {
	fmt.Println(item)
}
