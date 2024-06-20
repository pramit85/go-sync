/* The sync.Cond condition variable can be used to coordinate those goroutines that want to share resources. When the state of shared resources changes, it can be used to notify goroutines blocked by a mutex */

package main

import (
	"fmt"
	"sync"
	"time"
)

type employee struct {
	name string
}

var isWorkDone = false

func (e *employee) updateName(name string, c *sync.Cond) {

	c.L.Lock()
	e.name = name
	isWorkDone = true
	c.L.Unlock()
	fmt.Println("Employee name : ", e.name)

}

func (e *employee) getName(parm string, c *sync.Cond) {
	c.L.Lock()
	for !isWorkDone {
		c.Wait()
	}
	fmt.Println("reader ", parm, " Employee Name :", parm, e.name)

	c.L.Unlock()
}
func main() {

	var mu sync.Mutex

	cond := sync.NewCond(&mu)

	e := employee{}

	go e.getName("1 reader", cond)
	go e.getName("2 reader", cond)
	go e.getName("3 reader", cond)
	e.updateName("ram", cond)

	time.Sleep(4 * time.Second)
}
