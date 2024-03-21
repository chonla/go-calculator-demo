package main

import (
	"calculator/calculator"
	"fmt"
)

func main() {
	adder := calculator.NewAdder()
	calc := calculator.NewCalculator(adder)

	fmt.Println(calc.Add(4, 8))
}
