package main

import (
	"fmt"
	"strconv"
	"os"
)

func main() {
	
	for {
	var op string
	fmt.Print("Please select an operator: + - / * or 0 to exit: ")
	fmt.Scanln(&op)
if op=="0" {
	os.Exit(0)
	}
	var num1 string
	fmt.Printf("Input first number: ")
	fmt.Scanln(&num1)

	var num2 string
	fmt.Printf("Input second number: ")
	fmt.Scanln(&num2)
	if op=="0" {

	}else if op == "+" {
		firstNumber, _ := strconv.Atoi(num1)
		secondNumber,_ := strconv.Atoi(num2)
		result := firstNumber + secondNumber
		fmt.Printf("Result:%d\n", result)
	} else if op == "-" {
		firstNumber,_:= strconv.Atoi(num1)
		secondNumber,_ := strconv.Atoi(num2)
		result := firstNumber - secondNumber
		fmt.Printf("Result:%d\n", result)
	} else if op == "*" {
		firstNumber,_ := strconv.ParseFloat(num1, 64)
		secondNumber,_ := strconv.ParseFloat(num2, 64)
		result := firstNumber * secondNumber
		fmt.Printf("Result:%f\n", result)
	} else if op == "/" {
		firstNumber,_ := strconv.ParseFloat(num1, 64)
		secondNumber,_ := strconv.ParseFloat(num2, 64)
		result := firstNumber / secondNumber
		fmt.Printf("Result:%f\n", result)
	}
	}
}
