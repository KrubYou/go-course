package main

import "fmt"

var grade int = 0

func main() {

	fmt.Print("Enter your grade: ")
	fmt.Scanf("%d", &grade)

	if grade >= 80 {
		fmt.Println("A")
	} else if grade >= 70 {
		fmt.Println("B")
	} else if grade >= 60 {
		fmt.Println("C")
	} else if grade >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
