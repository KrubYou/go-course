package main

import "fmt"

func zerovalue(ivalue int) {
	ivalue = 0
}

func zeropointer(ipointer *int) {
	*ipointer = 0
}

func main() {
	i := 1
	fmt.Println("i=", i)

	zerovalue(i)
	fmt.Println("zerovalue=", i)

	zeropointer(&i)
	fmt.Println("zeropointer value =", i)
	fmt.Println("pointer address =", &i)
}
