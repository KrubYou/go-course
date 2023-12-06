package main

import "fmt"

var product = make(map[string]float64)

func main() {

	//add
	product["Coke"] = 12.5
	product["Pepsi"] = 13.5
	product["Fanta"] = 14.5
	product["Sprite"] = 15.5
	fmt.Println("product", product)

	//delete
	delete(product, "Coke")
	fmt.Println("product", product)

	//update
	product["Pepsi"] = 15000
	fmt.Println("product", product)

	value1 := product["Pepsi"]
	fmt.Println("value1", value1)

	courseName := map[string]string{"101": "Android", "102": "iOS", "103": "React Native", "104": "Go"}
	fmt.Println("courseName", courseName)

}
