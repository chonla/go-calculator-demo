package main

import (
	"calculator/calculator"
	"fmt"
)

func main() {
	calc := InitializeCalculator()

	fmt.Println(calc.Add(4, 8))
}

func InitializeCalculator() calculator.Calculator {
	adder := calculator.NewAdder()
	calc := calculator.NewCalculator(adder)

	return calc
}
