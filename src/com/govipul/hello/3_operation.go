package main

import (
	"fmt"
	"math"
)

func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

func main3() {
	var myInt8 int8 = 97

	var myInt = 1200
	var myUint uint = 500

	var myHexNumber = 0xFF  // Use prefix '0x' or '0X' for declaring hexadecimal numbers
	var myOctalNumber = 034 // Use prefix '0' for declaring octal numbers

	fmt.Printf("%d, %d, %d, %#x, %#o\n", myInt8, myInt, myUint, myHexNumber, myOctalNumber)

	num1, num2 := 5.6, 9.5

	fmt.Println(add(num1, num2))

	w1, w2 := "Hey", "Jeg"

	fmt.Println(multiple(w1, w2))

	//Numbers
	var a, b = 4, 5

	var res1 = (a + b) * (a + b) / 2

	a++

	b += 10

	var res2 = a ^ b

	var radius = 3.5

	var res3 = math.Pi * math.Pow(radius, 2)

	fmt.Printf("res1: %v, res2: %v, res3: %v\n", res1, res2, res3)

	//Boolean
	var truth = 3 <= 5
	var falsehood = 10 != 10

	var res4 = 10 > 20 && 5 == 5
	var res5 = 2*2 == 4 || 10%3 == 0

	fmt.Println(truth, falsehood, res4, res5)

	//Complex Number
	var c1 = 3 + 5i
	var c2 = 2 + 4i

	var res6 = c1 + c2
	var res7 = c1 - c2
	var res8 = c1 * c2
	var res9 = c1 / c2

	fmt.Println(res6, res7, res8, res9)

	// String
	var website = "\twww.google.com\t\n"
	var description = `
	\t\tSteve Jobs was an American entrepreneur and inventor.
	He was the CEO and co-founder of Apple Inc.\t\t\n`
	fmt.Println(website, description)

	//
}
