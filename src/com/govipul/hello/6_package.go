package main

import (
	"com/govipul/hello/numbers"
	"com/govipul/hello/strings"
	"com/govipul/hello/strings/greetings"
	"fmt"
	//"rsc.io/quote"
)

func main6() {

	fmt.Println(numbers.IsPrime(19))

	fmt.Println(strings.Reverse("TECHNIA"))

	fmt.Println(greetings.WelcomeText)

	//fmt.Println(quote.Go())
}
