package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	x := make([]int, 11)
	for i := 10; i >= 0; i-- {
		wg.Add(1)

		go func(i int) {
			m.Lock()
			defer m.Unlock()
			x[1] = i
			defer wg.Done()

			fmt.Println(i)
		}(i)
	}

	wg.Wait()
}
