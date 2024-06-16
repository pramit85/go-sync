/* An RWMutex is a reader/writer mutual exclusion lock. The lock can be held by an arbitrary number of readers or a single writer */

package main

import (
	"fmt"
	"sync"
)

type employee struct {
	salary int
	mu     sync.RWMutex
}

func (e *employee) updateSalary(incrementSalary int, wg *sync.WaitGroup) {
	defer wg.Done()
	e.mu.Lock()

	e.salary += incrementSalary

	fmt.Println("updated salary is :", e.salary)

	e.mu.Unlock()

}

func (e *employee) getSalary(wg *sync.WaitGroup) {
	defer wg.Done()
	e.mu.RLock()
	defer e.mu.RUnlock()
	fmt.Println("Now salary is :", e.salary)
}

func main() {

	var wg sync.WaitGroup

	wg.Add(6)

	e := employee{}

	go e.getSalary(&wg)
	go e.getSalary(&wg)
	go e.updateSalary(10, &wg)
	go e.getSalary(&wg)
	go e.getSalary(&wg)
	go e.getSalary(&wg)

	wg.Wait()
}
