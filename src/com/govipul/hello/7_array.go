package main

import "fmt"

func main() {
	var a [10]int64
	fmt.Println(a)

	var b [5]string
	fmt.Println(b)

	var c [5]complex128
	fmt.Println(c)

	var x [5]int

	x[0] = 100
	x[1] = 101
	//x[2] = 102
	x[3] = 103
	x[4] = 104
	fmt.Println(x)

	var y = [5]int{2, 4, 6, 8, 10}
	fmt.Println(y)

	aa := [...]int{3, 5, 7, 9}
	fmt.Println(aa)

	var a1 = [5]int{3, 6, 9, 12, 15}
	//var b1 [10]int = a1 //Illegal assignment as array of 5 cannot be assigned to Array of 10
	var b1 [5]int = a1
	fmt.Println(b1)

	//String arrays
	aa1 := [5]string{"English", "Japanese", "Spanish", "French", "Hindi"}
	aa2 := aa1
	aa2[2] = "German"
	fmt.Println(aa1) // The array `a1` remains unchanged
	fmt.Println(aa2)

	//Iteration
	a3 := [...]string{"Mark", "Bill", "Larry"}

	/*for i := 0; i < 3; i++ {
		fmt.Println(a3[i])
	}*/
	for i := 0; i < len(a3); i++ {
		fmt.Println(a3[i])
	}

	a4 := [4]float64{3.5, 7.2, 4.8, 9.5}
	sum := float64(0)

	for i := 0; i < len(a4); i++ {
		sum += a4[i]
	}
	fmt.Println(sum)

	sum1 := float64(0)
	/*
		for index, value := range a4 { // this is give error as index was declared but never used
			sum1 += value				// so solution is below
		}
	*/
	for _, value := range a4 {
		sum1 += value
	}
	fmt.Println(sum1)

	daysOfWeek := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	for index, value := range daysOfWeek {
		fmt.Printf("%d th is day of week %s\n", index+1, value)
	}

	//Multidimensional Array
	a5 := [2][2]int64{
		{3, 5},
		{10, 32},
	}
	fmt.Println(a5)

	a6 := [3][4]float64{
		{1, 3},
		{4.5, -3.1, 7.4, 2},
		{6, 2, 11},
	}
	fmt.Println(a6)
}
