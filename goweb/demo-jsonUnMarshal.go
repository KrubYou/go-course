package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	e := employee{}
	err := json.Unmarshal([]byte(`{"ID":1,"EmployeeName":"Puttipong","Tel":"0812345678","Email":"puttipong@mail.cc"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e.EmployeeName)
}
