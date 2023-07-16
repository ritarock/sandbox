package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func generator(ctx context.Context, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case out <- num:
			}
		}

		close(out)
		fmt.Println("generator closed")
	}()

	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	gen := generator(ctx, 1)

	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)
	}

	cancel()

	wg.Wait()

	part1()
	part2()
	part3()
	part4()
}

func part1() {
	ctx0 := context.Background()
	ctx1, _ := context.WithCancel(ctx0)

	// G1
	go func(ctx1 context.Context) {
		ctx2, cancel2 := context.WithCancel(ctx1)

		// G2-1
		go func(ctx2 context.Context) {
			// G2-2
			go func(ctx2 context.Context) {
				select {
				case <-ctx2.Done():
					fmt.Println("G2-2 canceled")
				}
			}(ctx2)

			select {
			case <-ctx2.Done():
				fmt.Println("G2-1 canceled")
			}

		}(ctx2)

		cancel2()

		select {
		case <-ctx1.Done():
			fmt.Println("G1 canceled")
		}

	}(ctx1)

	time.Sleep(time.Second)
}

func part2() {
	ctx0 := context.Background()

	ctx1, cancel1 := context.WithCancel(ctx0)

	// G1-1
	go func(ctx1 context.Context) {
		select {
		case <-ctx1.Done():
			fmt.Println("G1-1 canceled")
		}
	}(ctx1)

	// G1-2
	go func(ctx1 context.Context) {
		select {
		case <-ctx1.Done():
			fmt.Println("G1-2 canceled")
		}
	}(ctx1)

	cancel1()

	time.Sleep(time.Second)
}

func part3() {
	ctx0 := context.Background()
	ctx1, cancel1 := context.WithCancel(ctx0)

	// G1
	go func(ctx1 context.Context) {
		select {
		case <-ctx1.Done():
			fmt.Println("G1 canceled")
		}
	}(ctx1)

	ctx2, _ := context.WithCancel(ctx0)

	// G2
	go func(ctx2 context.Context) {
		select {
		case <-ctx2.Done():
			fmt.Println("G2 canceled")
		}
	}(ctx2)

	cancel1()

	time.Sleep(time.Second)
}

func part4() {
	ctx0 := context.Background()
	ctx1, _ := context.WithCancel(ctx0)

	// G1
	go func(ctx1 context.Context) {
		ctx2, cancel2 := context.WithCancel(ctx1)

		// G2
		go func(ctx2 context.Context) {
			ctx3, _ := context.WithCancel(ctx2)

			// G3
			go func(ctx3 context.Context) {
				select {
				case <-ctx3.Done():
					fmt.Println("G3 canceled")
				}
			}(ctx3)

			select {
			case <-ctx2.Done():
				fmt.Println("G2 canceled")
			}
		}(ctx2)

		cancel2()

		select {
		case <-ctx1.Done():
			fmt.Println("G1 canceled")
		}

	}(ctx1)

	time.Sleep(time.Second)
}
