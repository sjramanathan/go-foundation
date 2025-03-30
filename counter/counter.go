package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	count := int64(0)
	const n = 10
	var wg sync.WaitGroup

	wg.Add(n)
	for range n {
		go func() {
			defer wg.Done()
			for range 10_000 {
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
