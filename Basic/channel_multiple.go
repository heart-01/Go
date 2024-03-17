package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	quit := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- "End"
	}()

	task(ch, quit)
}

func task(ch, quit chan string) {
	for {
		select {
		case ch <- "Running...":
		case <-quit:
			fmt.Println("Quit")
			return
		default:
			fmt.Println("Waiting...")
			time.Sleep(1 * time.Second)
		}

	}
}
