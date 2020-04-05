package main

import "fmt"

func main8() {
	var s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t := []int{2, 4, 8, 16, 32, 64}
	fmt.Println("s = ", s)
	fmt.Println("t = ", t)

	var a = []string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}
	var a1 []string = a[1:4]
	fmt.Println("a = ", a)
	fmt.Println("a1 = ", a1)

	c := [5]string{"C", "C++", "Java", "Python", "Go"}
	fmt.Println("C1 =", c[1:4])
	fmt.Println("C2 =", c[:3])
	fmt.Println("C3 =", c[2:])
	fmt.Println("C4 =", c[:])

	cities := [9]string{"New York", "London", "Chicago", "Beijing", "Delhi", "Mumbai", "Bangalore", "Hyderabad", "Hong Kong"}

	asiancities := cities[3:]
	indianCities := asiancities[1:5]

	fmt.Println(asiancities)
	fmt.Println(indianCities)

	weeks := [7]string{"Mon", "Tue", "Wed", "Thru", "Fri", "Sat", "Sun"}

	slice1 := weeks[1:]
	slice2 := weeks[3:]

	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)

	slice1[0] = "TUE"
	slice1[1] = "WED"
	slice1[2] = "THRU"

	slice2[1] = "FRIDAY"
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)

	number := [6]int{1, 2, 3, 4, 5, 6}
	slice3 := number[1:4]

	fmt.Printf("s=%v, len = %d, cap = %d\n", slice3, len(slice3), cap(slice3))

	s1 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println("Original")
	fmt.Printf("s1=%d, len=%d, capacity=%d\n", s1, len(s1), cap(s1))

	s1 = s1[1:5]
	fmt.Println("Slice 1 --> 1-5")
	fmt.Printf("s1=%d, len=%d, capacity=%d\n", s1, len(s1), cap(s1))

	s1 = s1[:8]
	fmt.Println("Slice 2 --> 0-8")
	fmt.Printf("s1=%d, len=%d, capacity=%d\n", s1, len(s1), cap(s1))

	s1 = s1[2:]
	fmt.Println("Slice 3 --> 2-LEN")
	fmt.Printf("s1=%d, len=%d, capacity=%d\n", s1, len(s1), cap(s1))

	make1 := make([]int, 1, 10)
	fmt.Printf("s=%v, len=%d, cap=%d\n", make1, len(make1), cap(make1))

	var s2 []int
	fmt.Printf("s=%v, len=%d, cap=%d\n", s2, len(s2), cap(s2))

	if s2 == nil {
		fmt.Println("s2 is null")
	}

	//Copy
	src := []string{"Sublime", "VSCode", "IntelliJ", "Eclipse"}
	dest := make([]string, len(src)-2)
	copiedElement := copy(dest, src)
	fmt.Println("Source = ", src)
	fmt.Println("Destination = ", dest)
	fmt.Println("Copy = ", copiedElement)

	//append
	/*
		slice11 := []string{"C", "C++", "Java"}
		slice12 := append(slice11, "Python", "Ruby", "Go")
	*/
	slice11 := make([]string, 3, 10)
	copy(slice11, []string{"C", "C++", "Java"})
	slice12 := append(slice11, "Python", "Ruby", "Go")

	fmt.Println("Slice1 ", slice11)
	fmt.Println("Slice2 ", slice12)

	slice11[0] = "C#"
	fmt.Println("Slice1 ", slice11)
	fmt.Println("Slice2 ", slice12)

	//nil slice
	var sNil []string
	sNil = append(sNil, "Cat", "Dog", "Tiger")
	fmt.Println("sNil :", sNil)

	//one slice to another
	slice21 := []string{"Jack", "John", "Peter"}
	slice22 := []string{"Bill", "Mark", "Steve"}

	slice23 := append(slice21, slice22...)
	fmt.Println("slice23:", slice23)

	//Slice of Sclice
	scountry := [][]string{
		{"India", "China"},
		{"USA", "Canada"},
		{"Switzerland", "Germany"},
	}

	fmt.Println("Slice s = ", scountry)
	fmt.Println("length = ", len(scountry))
	fmt.Println("capacity = ", cap(scountry))

	//Loop in slice
	fmt.Println("--------------------")
	for i := 0; i < len(slice23); i++ {
		fmt.Println(slice23[i])
	}

	primeIndex := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	for index, number := range primeIndex {
		fmt.Printf("PrimeNumber(%d) : %d\n", index+1, number)
	}

	numbers := []float64{3.5, 7.4, 9.2, 5.4}

	sum := 0.0

	for _, number := range numbers {
		sum += number
	}
	fmt.Println("SUM = ", sum)
}
