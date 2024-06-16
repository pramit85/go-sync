/* A Mutex is a mutual exclusion lock that prevents other processes from entering a critical section of data
while a process occupies it to prevent race conditions from happening */

package main

import (
	"fmt"
	"sync"
)

type employee struct {
	salary int
	mu     sync.Mutex
}

func (e *employee) update(incrementSalary int, wg *sync.WaitGroup) {

	e.mu.Lock()
	defer wg.Done()

	fmt.Println("Increment Salary : ", incrementSalary)

	e.salary += incrementSalary

	fmt.Println("Now Salary : ", e.salary)

	e.mu.Unlock()

}

func main() {
	var wg sync.WaitGroup
	e := employee{}
	wg.Add(3)
	go e.update(10, &wg)
	go e.update(15, &wg)
	go e.update(20, &wg)

	wg.Wait()

}
