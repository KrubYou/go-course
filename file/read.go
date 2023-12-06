package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../currency.csv")
	if err != nil {
		panic(err)
	}

	count := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		item := strings.Split(line, ",")
		//fmt.Printf("Line %d : %s\n ", count, line)
		fmt.Println(item[3])
		count++
	}
}
