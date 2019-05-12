package main

import (
	employeeDomain "github.com/sangianpatrick/ss-golang/oop_basic3/modules/employee/domain"
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

	// createEmployee(eD)
	printAllEmployee(eD)

}
