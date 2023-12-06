package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {
	employeeList := []employee{}
	employee1 := employee{employeeID: "001", employeeName: "Puttipong", phone: "0812345678"}
	employee2 := employee{employeeID: "002", employeeName: "Pattaraporn", phone: "0812345678"}
	employee3 := employee{employeeID: "003", employeeName: "Pattaraporn", phone: "0812345678"}

	employeeList = append(employeeList, employee1)
	employeeList = append(employeeList, employee2)
	employeeList = append(employeeList, employee3)

	fmt.Println("emplist = ", employeeList)

}
