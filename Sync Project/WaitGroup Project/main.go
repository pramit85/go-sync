/*A WaitGroup waits for a collection of goroutines to finish*/

package main

import (
	"fmt"
	"sync"
)

func callmethod(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("calling method...")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go callmethod(&wg)

	wg.Wait()
}
