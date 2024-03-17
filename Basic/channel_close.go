package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	go routine(ch, 10)

	for true {
		value, ok := <-ch
		if !ok {
			fmt.Println("No more data")
			break
		}
		fmt.Println(value)
	}

	/* Refactor code using range
	for value := range ch {
		fmt.Println(value)
	}
	*/
}

func routine(c chan int, countTo int) {
	for i := 0; i < countTo; i++ {
		c <- i
	}
	close(c)
}
