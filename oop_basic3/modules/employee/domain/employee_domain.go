package domain

import (
	"github.com/sangianpatrick/ss-golang/oop_basic3/modules/employee"
	"github.com/sangianpatrick/ss-golang/oop_basic3/modules/employee/model"
)

// EmployeeDomain contains employee behavior and entity
type EmployeeDomain struct {
	mr employee.MongoRepository
}

// NewEmployeeDomain acts like constructor
func NewEmployeeDomain(mr employee.MongoRepository) employee.Domain {
	return &EmployeeDomain{
		mr: mr,
	}
}

// CreateOne will create a new employee
func (ed *EmployeeDomain) CreateOne(employee model.Employee) error {
	return ed.mr.InsertOne(employee)
}

// GetByID return employee with spesific ID
func (ed *EmployeeDomain) GetByID(id string) (*model.Employee, error) {
	employee, err := ed.mr.FindByID(id)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

// GetByEmail return employee with spesific email
func (ed *EmployeeDomain) GetByEmail(id string) (*model.Employee, error) {
	employee, err := ed.mr.FindByEmail(id)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

// GetAll return all employee
func (ed *EmployeeDomain) GetAll() (*model.Employees, error) {
	employees, err := ed.mr.FindAll()
	if err != nil {
		return nil, err
	}
	return employees, nil
}
