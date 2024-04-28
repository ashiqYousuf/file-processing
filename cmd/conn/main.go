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

func work(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("working...")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	work(&wg)

	wg.Wait()
}
