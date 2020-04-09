package main

import "fmt"

func main10() {
	var x = 100
	var p *int = &x
	fmt.Println(*p)

	*p = 200
	fmt.Println("x=", x)
	fmt.Println("p=", p)
	fmt.Println("*p=", *p)

	//pointer to pointer
	fmt.Println("*******************")
	a1 := 100
	var p1 = &a1
	var p2 = &p1
	fmt.Println("a1=", a1)
	fmt.Println("p1=", p1)
	fmt.Println("*p1=", *p1)
	fmt.Println("p2=", p2)
	fmt.Println("*p2=", *p2)
	fmt.Println("**p2=", **p2)

	fmt.Println("*******************")
	var b1 = 200
	var p3 = &b1
	var p4 = &b1

	if p3 == p4 {
		fmt.Println("Both the address is pointing to same location", p3, p4)
	}
}
