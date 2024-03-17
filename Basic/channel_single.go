package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go routine(ch, 1)
	go routine(ch, 2)
	go routine(ch, 3)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	time.Sleep(1 * time.Second)
}

func routine(ch chan int, payload int) {
	ch <- payload
}
