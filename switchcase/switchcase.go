package main

import "fmt"

func main() {
	input := "red"
	fmt.Scanf("%s", &input)
	switch input {
	case "blue":
		println("input is #0000FF")
	case "red":
		println("input is #FF0000")
	case "green":
		println("input is #00FF00")
	case "yellow":
		println("input is #FFFF00")
	default:
		println("input is not a color")
	}
}
