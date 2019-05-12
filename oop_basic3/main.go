package main

import (
	"encoding/json"
	"fmt"

	"github.com/sangianpatrick/ss-golang/oop_basic3/modules/employee"
	employeeDomain "github.com/sangianpatrick/ss-golang/oop_basic3/modules/employee/domain"
	"github.com/sangianpatrick/ss-golang/oop_basic3/modules/employee/model"
	employeeRepository "github.com/sangianpatrick/ss-golang/oop_basic3/modules/employee/repository"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	// pool connection to mongodb
	mgoSession, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:     []string{"localhost:27017"},
		Username:  "patrick",
		Password:  "14qwafzx",
		Database:  "gotest",
		PoolLimit: 5,
	})

	if err != nil {
		panic(err)
	}

	//-------------------------------------------------------------------------
	// Mounting MongoDB to Employee Repository
	// and pass it to Employee Domain
	//-------------------------------------------------------------------------
	eMR, err := employeeRepository.NewEmployeeMongoRepository(mgoSession)
	if err != nil {
		panic(err)
	}
	eD := employeeDomain.NewEmployeeDomain(eMR)

	//--------------------------------------------------------------------------
	// Go to employee handler function (CRUD)
	//--------------------------------------------------------------------------
	// var person = model.Employee{
	// 	ID:        "EMP0002",
	// 	FirstName: "Patrick",
	// 	LastName:  "Sangian",
	// 	Email:     "patricksangian@email.com",
	// 	Phone:     "093459345354",
	// 	Address:   "Jl. Karet Karya 1 No. 5",
	// 	Position:  "Backend Developer",
	// }
	// createEmployee(person, eD)
	// printEmployeeByID("EMP0001", eD)
	printAllEmployee(eD)
}

//----------------------------------------------------------------------------
// List of Handler function
//----------------------------------------------------------------------------

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
		employeeByte, _ := json.Marshal(employee)
		employeeJSON := string(employeeByte)
		fmt.Printf("Print Employee By ID: \n%v\n", employeeJSON)
	}
}

func printEmployeeByEmail(email string, d employee.Domain) {
	employee, err := d.GetByEmail(email)
	if err != nil {
		fmt.Println(err)
	} else {
		employeeByte, _ := json.Marshal(employee)
		employeeJSON := string(employeeByte)
		fmt.Printf("Print Employee By Email: \n%v\n", employeeJSON)
	}
}

func printAllEmployee(d employee.Domain) {
	employees, err := d.GetAll()
	if err != nil {
		fmt.Println(err)
	} else {
		employeesByte, _ := json.Marshal(employees)
		employeesJSON := string(employeesByte)
		fmt.Printf("Print Employee By ID: \n%v\n", employeesJSON)
	}
}
