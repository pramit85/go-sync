package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map

	//number of store and load goroutines
	numStoreGoroutines := 3
	numLoadGoroutines := 3

	var wg sync.WaitGroup
	wg.Add(numStoreGoroutines)

	//store goroutines
	for i := 0; i < numStoreGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			value := id + 5
			sm.Store(key, value)
			fmt.Printf("Storeed %s with value %d\n", key, value)
		}(i)
	}

	var loadWg sync.WaitGroup
	loadWg.Add(numLoadGoroutines)

	//store goroutines
	for i := 0; i < numLoadGoroutines; i++ {
		go func(id int) {
			defer loadWg.Done()
			key := fmt.Sprintf("key%d", id)
			value, ok := sm.Load(key)
			if ok {
				fmt.Printf("Loaded %s with value %d\n", key, value)
			} else {
				fmt.Printf("key %s not found \n", key)
			}
		}(i)
	}

	wg.Wait()
	loadWg.Wait()

	//Example of iterating over all entries in sync.Map
	finalState := make(map[string]int)
	sm.Range(func(key, value interface{}) bool {
		finalState[key.(string)] = value.(int)
		return true
	})

	fmt.Println("final map state (with sync.Map):", finalState)
}
