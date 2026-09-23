package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func main() {

	employees := []Employee{
		{ID: 1, Name: "Ray", Salary: 50000},
		{ID: 2, Name: "John", Salary: 60000},
		{ID: 3, Name: "Sam", Salary: 55000},
	}

	fmt.Println("Employee List:")

	for _, employee := range employees {
		fmt.Println(employee)

	}
	fmt.Println("Employee Management Application")
}