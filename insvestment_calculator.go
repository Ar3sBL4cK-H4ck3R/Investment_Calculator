package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64
	expectedReturnRate := 4.5
	var years float64

	println("\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	println("############### Investment Calculator ###############")
	println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$\n")

	fmt.Print("Enter Investment Ammount $: ")
	fmt.Scanln(&investmentAmount)

	fmt.Print("Enter Expected Return Rate: ")
	fmt.Scanln(&expectedReturnRate)

	fmt.Print("Enter Investment Horizon: ")
	fmt.Scanln(&years)

	futureValue, futureRealValue := calculations(investmentAmount, expectedReturnRate, years)

	output("Future Value $: %.2f\n", futureValue)
	fmt.Printf("Future Real Value $: %.2f\n", futureRealValue)
}

func output(value1 string, value2 float64) {
	fmt.Printf(value1, value2)
}

func calculations(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
}
