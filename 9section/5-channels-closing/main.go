package main

import (
	"fmt"
	"sync"
)

func main() {

	jobs := make(chan int, 5)
	var wg sync.WaitGroup
	// done := make(chan bool)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			r, ok := <-jobs
			if ok {
				fmt.Println("Got this message", r)
			} else {
				fmt.Println("Channel closed", r)
				//done <- true
				return
			}
		}
	}(&wg)

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("Sending ", i)
	}

	close(jobs)

	// <-done another way to wait group, for example for concurent processing
	wg.Wait()

}
