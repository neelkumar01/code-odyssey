package main

import "fmt"

type employee interface{
	getName() 	 string
	getSalary() int
}

type partTimeEmployee struct{
	name 		 string
	hourlyPay 	 int
	hoursPerYear int
}
type fullTimeEmployee struct{
	name   string
	salary int
}

func (pte partTimeEmployee) getName() string {
	return pte.name
}
func (fte fullTimeEmployee) getName() string {
	return fte.name
}
func (pte partTimeEmployee) getSalary() int {
	return pte.hourlyPay * pte.hoursPerYear
}
func (fte fullTimeEmployee) getSalary() int {
	return fte.salary
}

func getInfo(e employee) {
	fmt.Println(e.getName(), e.getSalary())
}

func main() {
	getInfo(partTimeEmployee{
		name: "neel",
		hourlyPay: 500,
		hoursPerYear: 80,
	})

	getInfo(fullTimeEmployee{
		name: "Ram",
		salary: 204000,
	})
}
