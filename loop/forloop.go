package main

import "fmt"

const count int = 10

func main() {
	// i := 0

	// for i < 10 {
	// 	println("i=", i)
	// 	i++

	// }

	// for j := 0; j < count; j++ {
	// 	println("j=", j)
	// }

	var input string
	for {
		fmt.Scanf("%s", &input)
		fmt.Println("input is ", input)
		if input == "stop" {
			break
		}
	}
}
