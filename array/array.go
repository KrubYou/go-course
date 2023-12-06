package main

import "fmt"

var productName [4]string

func main() {

	productName[0] = "Coke"
	productName[1] = "Pepsi"
	productName[2] = "Fanta"
	productName[3] = "Sprite"

	price := [4]float32{12.5, 13.5, 14.5, 15.5}

	fmt.Println(productName)
	fmt.Println(price)
}
