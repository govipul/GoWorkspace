package main

import "fmt"

func main() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	var p1 = Person{"Vipul", "Goswami", 32}
	fmt.Println("P1:", p1)

	var p2 = Person{FirstName: "Amrita", LastName: "Goswami", Age: 32}
	fmt.Println("P2:", p2)

	p2.Age = 30
	fmt.Println("P2, V2:", p2)

	ps := &p2
	fmt.Println("PS:", ps)
	fmt.Println("*PS:", *ps)

	ps.Age = 31
	fmt.Println("PS:", ps)
	fmt.Println("*PS:", *ps)

	type Point struct {
		X float64
		Y float64
	}

	po1 :=
		Point{10, 20}
	po2 := po1

	fmt.Println(po1)
	fmt.Println(po2)

	po3 := Point{10, 20}

	if po1 == po3 {
		fmt.Printf("PO1: %v and PO3: %v are not equal\n", po1, po3)
	}
	if po2 == po1 {
		fmt.Printf("%v and %v are equal\n", po1, po2)
	}
}
