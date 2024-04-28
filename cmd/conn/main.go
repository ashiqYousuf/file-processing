package main

import (
	"fmt"
	"sync"
)

/*
Concurrency is about dealing with lots of things at once.
Parallelism is about doing lots of things at once.
 - Rob Pike
*/

func Incr(n *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	// atomic.AddInt64(n, 1)
	*n += 1
}

func main() {
	var n int64

	var wg sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go Incr(&n, &wg)
	}

	wg.Wait()

	fmt.Println("value", n)
}
