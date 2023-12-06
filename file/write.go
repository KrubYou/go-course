package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	data1 := []byte("Hi\n Puttipong!")
	err := os.WriteFile("data.txt", data1, 0644)
	check(err)
	f, err := os.Create("employeeName")
	check(err)
	defer f.Close()

	data2 := []byte("Puttipong\n Huangnam")
	os.WriteFile("employeeName.txt", data2, 0644)

}
