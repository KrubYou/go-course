package main

import "fmt"

func hello() {
	fmt.Println("Hello, Puttipong!")
}

func plus(a, b int) int {
	return a + b
}

func main() {
	hello()
	fmt.Println("result = ", plus(10, 20))

}
