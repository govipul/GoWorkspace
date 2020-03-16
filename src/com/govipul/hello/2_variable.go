package main

import (
	"fmt"
	"math"
)

func main2() {
	var firstName string = "Vipul"
	var lastName string = "Goswami"
	var age int = 32

	var (
		salary           = 770000.0
		isConfirmed bool = true
	)

	fmt.Printf("firstName: %s, lastName: %s, Age: %d, Salary: %f, IsConfirmed: %t\n",
		firstName, lastName, age, salary, isConfirmed)

	name := "Vipul Goswami"
	nAge, nSalary, isProgram := 32, 1500000, true

	fmt.Println(name, nAge, nSalary, isProgram)

	fmt.Println(math.Abs(-2.78556))

}
