package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine") // Starting a go routine
	fmt.Println("main")

	for i := 0; i < 3; i++ {
		// BUG: All goroutines use the "i" for the for loop

		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(10 * time.Millisecond)

	ch := make(chan string)

	go func() {
		ch <- "hi" // send
	}()
	msg := <-ch // receive
	fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("message #%d", i+1)
			ch <- msg
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("got: ", msg)
	}

	msg, ok := <-ch
	fmt.Printf("closed: %#v (ok=%v)\n", msg, ok) // receive from a closed channel will return the zero value without blocking
	fmt.Println()

	sleepSort([]int{4, 5, 2, 3, 1})
}

/*
	For every value "n" in values, spin a goroutine that will
	- sleep "n" milliseconds
	- send "n" over a channel

	In the function body,
	collecet values from the channel to a slice and return it.
*/

func sleepSort(values []int) []int {
	if len(values) == 0 {
		return nil
	}

	ch := make(chan int)

	for _, v := range values {
		go func() {
			time.Sleep(time.Duration(v) * time.Millisecond)
			ch <- v
		}()
	}

	sortedValues := make([]int, 0)
	for range values {
		msg, ok := <-ch
		if !ok {
			break
		}

		sortedValues = append(sortedValues, msg)
	}

	fmt.Printf("%-20s %#v\n", "original values:", values)
	fmt.Printf("%-20s %#v\n", "sorted values:", sortedValues)
	return sortedValues
}
