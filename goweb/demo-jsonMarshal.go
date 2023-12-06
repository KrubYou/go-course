package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {

	data, _ := json.Marshal(&employee{ID: 1, EmployeeName: "Puttipong", Tel: "0812345678", Email: "puttipong@mail.comm"})
	fmt.Println(string(data))
}
