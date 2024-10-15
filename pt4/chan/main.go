package main

import (
	"fmt"
	"sync"
)

func worker(i int, c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for p := range c {
		fmt.Printf("Worker id: %d: payload: %d\n", i, p)
	}
}

func main() {
	c := make(chan int)

	wg := new(sync.WaitGroup)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go worker(i, c, wg)
	}

	for j := 10; j > 0; j-- {
		c <- j
	}

	close(c)

	wg.Wait()
}
