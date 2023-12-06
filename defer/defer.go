package main

import "fmt"

func add(value1, value2 float32) float32 {
	return value1 + value2
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i=", i)
	}
}

func deferloop() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("i=", i)
	}
}

func main() {

	// defer fmt.Println("Hello, Puttipong!")
	// defer fmt.Println("result = ", add(10, 20))
	// a := add(1.1, 2.2)
	// fmt.Println(a)
	loop()
	deferloop()

}
