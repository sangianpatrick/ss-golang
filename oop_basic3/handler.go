package main

import (
	"fmt"

	"github.com/sangianpatrick/ss-golang/oop_basic3/modules/employee"
	"github.com/sangianpatrick/ss-golang/oop_basic3/modules/employee/model"
)

func createEmployee(employee model.Employee, d employee.Domain) {
	err := d.CreateOne(employee)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("New Employee has been successfully added")
	}
}

func printEmployeeByID(id string, d employee.Domain) {
	employee, err := d.GetByID(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Print Employee By ID: %v\n", employee)
	}
}

func printEmployeeByEmail(email string, d employee.Domain) {
	employee, err := d.GetByEmail(email)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Print Employee By Email: %v\n", employee)
	}
}

func printAllEmployee(d employee.Domain) {
	employees, err := d.GetAll()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Print Employee By Email: %v\n", employees)
	}
}
