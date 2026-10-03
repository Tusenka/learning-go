package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(message string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	fmt.Println("sayHello", message)
}

func main() {

	var wg sync.WaitGroup

	/*
	 1. Add outside of the goroutine
	 2. You must decrease the counter by calling wg.Done inside the goroutine
	 3. Do not forget to call wg.Wait()
	 4. Always pass a reference/pointer of the wait group variable instead of a copy
	*/

	totalJobs := 5

	for i := range totalJobs {
		wg.Add(1)
		go sayHello(fmt.Sprintf("JOB %d", i), time.Second, &wg)
		//go sayHello(fmt.Sprintf("JOB %d", i), time.Second, wg)
		// wg.Done() //do not to that!

	}
	fmt.Println("Hello from Main() Goroutine")

	wg.Wait()

}
