package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func main() {
	employee := Employee{
		ID:     1,
		Name:   "Ray",
		Salary: 50000,
	}

	fmt.Println("Employee:", employee)
}
