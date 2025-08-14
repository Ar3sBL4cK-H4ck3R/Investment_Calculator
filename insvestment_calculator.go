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

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	futureValue, futureRealValue := calculations(investmentAmount, expectedReturnRate, years)

	output("Future Value $: %.2f\n", futureValue)
	fmt.Printf("Future Real Value $: %.2f\n", futureRealValue)

	// *We Can Also Store these string outputs in a variable as below:

	// string_outputFV := fmt.Sprintf("Future Value $: %.2f\n", futureValue)
	// string_outputFRV := fmt.Sprintf("Future Real Value $: %.2f\n", futureRealValue)

	// fmt.Print(string_outputFV, string_outputFRV)

	/* also we can use fmt.Printf() func to print all outputs in one fmt.printf() func
	by separating values by commmas (,) */

	// fmt.Printf("Future Value $: %.2f\nFuture Real Value $: %.2f\n", futureValue, futureRealValue)

	// we can also use backticks to use multiple lines in a String without getting an error
	//  if the string is much longer

	/* fmt.Printf(`Future Value $: %.2f\n
	Future Real Value $: %.2f\n`, futureValue, futureRealValue)*/

}

func output(value1 string, value2 float64) {
	fmt.Printf(value1, value2)
}

func calculations(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
	// return
}
