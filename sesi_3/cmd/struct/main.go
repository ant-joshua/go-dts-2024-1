package main

import "fmt"

type Employee struct {
	Division string
	Salary   int
	Person   Person
	Tax      int
}

func (e *Employee) CalculateTax(taxValue int) {
	e.Tax = e.Salary / taxValue
}

func (e *Employee) TotalSalaryWithTax() int {
	return e.Salary - e.Tax
}

type Person struct {
	Name    string
	Age     int
	Address *string
}

func main() {
	// var employee2 *Employee

	var employee = Employee{
		Division: "IT",
		Salary:   1000,
		Person: Person{
			Name: "Joshua",
			Age:  25,
		},
	}

	employee.CalculateTax(10)

	fmt.Printf("Total Salary with tax %d ", employee.TotalSalaryWithTax())

	// employee2 = &employee

	// println("Name: ", employee.Person.Name)
	// println("Name: ", employee2.Person.Name)

	// println(strings.Repeat("#", 10))

	// employee2.Person.Name = "Joshua Antonius"

	// println("Name: ", employee.Person.Name)
	// println("Name: ", employee2.Person.Name)

}
