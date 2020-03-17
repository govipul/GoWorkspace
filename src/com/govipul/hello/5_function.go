package main

import (
	"errors"
	"fmt"
	"math"
)

func sayHello() {
	fmt.Println("Hello World")
}

func avg(x, y float64) float64 {
	return (x + y) / 2
}

func getPriceChange(prevPrice, currentPrice float64) (float64, float64, error) {

	if prevPrice <= 0 {
		err := errors.New("Previous price cannot be 0 (zero)")
		return 0, 0, err
	}

	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange, nil
}

//Naked return
func getStockPriceChange(prevPrice, currentPrice float64) (change, percentChange float64) {
	change = currentPrice - prevPrice
	percentChange = (change / prevPrice) * 100
	return
}

func main() {

	sayHello()
	x, y := 5.75, 13.44
	fmt.Println(avg(x, y))

	//prevStockPrice := 75000.0
	prevStockPrice := 0.0
	currentStockPrice := 100000.0

	change, changePercent, err := getPriceChange(prevStockPrice, currentStockPrice)

	if err != nil {
		fmt.Println("There was en error: ", err)
	} else {
		if change < 0 {
			fmt.Printf("the stock price has be dropped by %.2f which is change of %.2f\n", math.Abs(changePercent), math.Abs(change))
		} else {
			fmt.Printf("the stock price has be raised by %.2f %% which is change of %.2f\n", changePercent, change)
		}
	}

	//Naked return
	fmt.Println("Naked Return:")
	prevStockPrice1 := 10000.0
	currentStockPrice1 := 9000.0
	change1, changePercent1 := getStockPriceChange(prevStockPrice1, currentStockPrice1)

	if change1 < 0 {
		fmt.Printf("the stock price has be dropped by %.2f %% which is change of %.2f\n", math.Abs(changePercent1), math.Abs(change1))
	} else {
		fmt.Printf("the stock price has be raised by %.2f %% which is change of %.2f\n", changePercent1, change1)
	}

	// Blank Identifier
	fmt.Println("Blank Identifier:")
	change2, _ := getStockPriceChange(prevStockPrice1, currentStockPrice1)

	if change2 < 0 {
		fmt.Printf("the stock price has be dropped %.2f\n", math.Abs(change2))
	} else {
		fmt.Printf("the stock price has be raised %.2f\n", change2)
	}
}
