package main

import "fmt"

var numberInt, numberInt2 int = 1000, 2000
var msg string = "Hello, Puttipong!"

func main() {
	numberfloat := 100.53
	fmt.Println("Hello, Puttipong!")
	fmt.Println("Number is", numberInt)
	fmt.Println("Number is", numberInt2)
	fmt.Println("Number is", numberfloat)
	fmt.Println("Message is", msg)
	fmt.Println(numberfloat + float64(numberInt))
	fmt.Println(msg + "hello")
	fmt.Println("my number is ", numberInt, "and", numberInt2)
}
