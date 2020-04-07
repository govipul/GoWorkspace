package main

import "fmt"

func main9() {

	var m = make(map[string]int)
	//Or can be initialized using curly braces
	//var m = map[string]int{}

	fmt.Println(m)

	if m == nil {
		fmt.Print("m is nil")
	}

	m["One Hundered"] = 100
	fmt.Println(m)

	//Assign with literals
	n := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println("Literal map:", n)
	fmt.Println("Literal map second value:", n["two"])

	//changing value
	tinderMatch := map[string]string{
		"Rajeev": "Angelia",
		"James":  "Sophia",
		"David":  "Emma",
	}

	fmt.Println("Tinder Match:", tinderMatch)

	tinderMatch["Rajeev"] = "Jennifer"
	fmt.Println("Tinder Match 2:", tinderMatch)

	fmt.Println("Tinder Match for Jack:", tinderMatch["Jack"])

	//How to check key is present in MAP
	value, ok := tinderMatch["Jack"]
	if !ok {
		fmt.Println("Tinder Match for jack not present:", value)
	}

	_, ok1 := tinderMatch["Jack"]
	if !ok1 {
		fmt.Println("Tinder Match for jack not present without value")
	}

	//Deleteing
	fileExtension := map[string]string{
		"Python": ".py",
		"C++":    ".cpp",
		"Java":   ".java",
		"Golang": ".go",
		"Kotlin": ".kt",
	}

	fmt.Println(fileExtension)

	delete(fileExtension, "Kotlin")
	fmt.Println(fileExtension)

	//delete key which doesnt exists
	delete(fileExtension, "Javascript")
	fmt.Println(fileExtension)

	//Maps are pass by refrence
	m2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	m3 := m2
	fmt.Println("M2: ", m2)
	fmt.Println("M3: ", m3)

	m3["Four"] = 4
	fmt.Println("M2: ", m2)
	fmt.Println("M3: ", m3)

	for file, extension := range fileExtension {
		fmt.Printf("%s file has following extension: %s \n", file, extension)
	}
}
