package main

import (
	"fmt"
	"time"
)

func process(num int, str string) {
	for i := 0; i <= num; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i, str)
	}
}

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	ch4 := make(chan bool)

	fmt.Println("start")

	go func() {
		go func() {
			process(2, "A")
			ch3 <- true
		}()

		go func() {
			process(2, "C")
			ch4 <- true
		}()
		<-ch3
		<-ch4
		ch1 <- true
	}()

	go func() {
		process(2, "B")
		ch2 <- true
	}()

	<-ch1
	<-ch2

	fmt.Println("Finish")
}
