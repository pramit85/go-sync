package main

import (
	"fmt"
	"sync"
)

// Employee represent an employee structure
type Employee struct {
	ID   int
	Name string
}

// create pool of Employee which can be reused.
var employeePool = &sync.Pool{
	New: func() interface{} {
		fmt.Println("Allocating new Employee")
		return &[]Employee{}
	},
}

func main() {
	//acquire a slice from the pool
	employees := employeePool.Get().(*[]Employee)

	//populate the employee list
	*employees = []Employee{
		{ID: 1, Name: "Ram"},
		{ID: 2, Name: "Rahim"},
		{ID: 3, Name: "Grumit"},
	}

	//do something with the employee list
	fmt.Println("Initial employee list", *employees)

	//Release the slice back to the pool (reset it)
	*employees = (*employees)[:0] //reset slice to zero length
	employeePool.Put(employees)

	//acquire another slice from the pool
	employees = employeePool.Get().(*[]Employee)

	//populate the employee list again
	*employees = []Employee{
		{ID: 1, Name: "ramesh"},
		{ID: 2, Name: "rohit"},
	}

	//do something with the new employee list
	fmt.Println("Initial employee list", *employees)

	//Release the second slice back to the pool
	*employees = (*employees)[:0] //reset slicce to zero length
	employeePool.Put(employees)

	//the pool will now have two empty slice that can be reused

	//example of resuing the pool.
	employees = employeePool.Get().(*[]Employee)

	//Example of reusing the pool:
	employees = employeePool.Get().(*[]Employee)
	fmt.Println("Resused empty employee list :", *employees)

	//Ensure the pool is cleared when done
	employeePool.Put(nil)

}
