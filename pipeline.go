package main

import "fmt"

func Producer(done <-chan struct{}) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-done:
				fmt.Println("Producer done")
				return
			case c <- i:
			}
		}
		close(c)
	}()
	return c
}

func Middleware(done <-chan struct{}, c <-chan int) <-chan int {
	mc := make(chan int)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Middleware done")
				return
			case n, ok := <-c:
				if !ok {
					close(mc)
					return
				}
				if n%2 == 0 {
					mc <- n
				} else {
					mc <- n * 2
				}
			}
		}
	}()
	return mc
}

func Consumer(done <-chan struct{}, c <-chan int) {
	for {
		select {
		case <-done:
			fmt.Println("Consumer done")
			return
		case n, ok := <-c:
			if !ok {
				return
			}
			fmt.Printf("%d\n", n)
		}
	}
}
