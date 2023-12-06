package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must number only", promt)
		panic(message)
	}

	return value
}

func getOperator() string {
	fmt.Printf("operator (+,-,*,/) = ")
	operator, _ := reader.ReadString('\n')
	return strings.TrimSpace(operator)
}

func add(value1 float64, value2 float64) float64 {
	return value1 + value2
}

func minus(value1 float64, value2 float64) float64 {
	return value1 - value2
}

func muliply(value1 float64, value2 float64) float64 {
	return value1 * value2
}

func divide(value1 float64, value2 float64) float64 {
	return value1 / value2
}

func main() {

	value1 := getInput(" value 1 = ")
	value2 := getInput(" value 2 = ")

	var result float64

	switch operator := getOperator(); operator {
	case "+":
		result = add(value1, value2)
	case "-":
		result = minus(value1, value2)
	case "*":
		result = muliply(value1, value2)
	case "/":
		result = divide(value1, value2)
	default:
		panic("unknown operator")

	}

	fmt.Printf("result = %v", result)
}
