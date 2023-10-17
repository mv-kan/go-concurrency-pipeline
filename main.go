package main

import (
	"fmt"
	"time"
)

func WithDone() {
	done := make(chan struct{})
	c := Producer(done)
	c = Middleware(done, c)
	go Consumer(done, c)
	close(done)
	time.Sleep(time.Second)
}

func WithoutDone() {
	c := Producer(nil)
	c = Middleware(nil, c)
	Consumer(nil, c)
}

func main() {
	fmt.Println("Without Done: ")
	WithoutDone()
	fmt.Println("With Done: ")
	WithDone()
}
