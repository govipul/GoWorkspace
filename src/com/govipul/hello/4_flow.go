package main

import "fmt"

func main4() {
	var x = 25

	if x%5 == 0 {
		fmt.Printf("%d is a multiple of 5\n", x)
	}

	var age = 21

	if age >= 17 && age <= 30 {
		fmt.Printf("Age is between 17 and 30 : %d\n", age)
	}

	var myage = 32

	if myage >= 18 {
		fmt.Println("You are eligible to work")
	} else {
		fmt.Println("You are not eligible to work")
	}

	const BMI = 21.0

	if BMI < 18.5 {
		fmt.Println("Underweight")
	} else if BMI >= 18.5 && BMI < 25.0 {
		fmt.Println("Your weight is normal")
	} else if BMI >= 25.0 && BMI < 30.0 {
		fmt.Println("Overweight")
	} else {
		fmt.Println("Thats enough, please control")
	}

	if n := 10; n%2 == 0 {
		fmt.Println("Given number is even")
	}

	if n := 15; n%2 == 0 {
		fmt.Printf("Given %d number is even\n", n)
	} else {
		fmt.Printf("Given %d number is odd\n", n)
	}

	//SWITCH
	//dayOfWeek := 4

	switch dayOfWeek := 5; dayOfWeek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thrusday")
	case 5:
		fmt.Println("Friday")
	case 6:
		{
			fmt.Println("Saturday")
			fmt.Println("Yaar weekend")
		}
	case 7:
		{
			fmt.Println("Sunday")
			fmt.Println("Yaar weekend")
		}
	}

	switch dayOfWeek := 8; dayOfWeek {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekdays")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid Day")
	}

	//Switch with no expression
	var BMIindex = 21.0
	switch {
	case BMIindex < 18.5:
		fmt.Println("Underweight")
	case BMIindex >= 18.5 && BMIindex < 25.0:
		fmt.Println("Your weight is normal")
	case BMIindex >= 25.0 && BMIindex < 30:
		fmt.Println("You are over weight")
	default:
		fmt.Println("Cmon its too much")
	}

	// FOR Loop
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	index := 0
	for ; index <= 2; index += 2 {
		fmt.Printf("%d ", index)
	}

	fmt.Println()

	// FOR loop similar to WHILE
	index2 := 2
	for index2 <= 20 {
		fmt.Printf("%d ", index2)
		index2 *= 2
	}

	//BREAK Statement
	// Infinite loop
	index3 := 1
	for {

		fmt.Printf("%d ", index3)

		if index3 > 30 {
			fmt.Println("Reached max value")
			break
		}

		index3 *= 2
	}

	//Continue Statement
	fmt.Println("Continue: ")
	for num := 1; num <= 10; num++ {
		if num%2 == 0 {
			continue
		}
		fmt.Printf("%d ", num)
	}
}
