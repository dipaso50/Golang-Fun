package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 1)
	var sl []int

	var wg sync.WaitGroup

	for x := 0; x < 3; x++ {
		//three go routines

		wg.Add(1)

		go func() {
			defer wg.Done()
			// ten integers
			for i := 0; i < 10; i++ {
				ch <- i
			}
		}()
	}

	//wait for all goroutines and close the channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	for l := range ch {
		sl = append(sl, l)
	}

	fmt.Printf("Slice len is %d \n", len(sl))
}
