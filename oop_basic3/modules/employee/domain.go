package employee

import (
	"github.com/sangianpatrick/ss-golang/oop_basic3/modules/employee/model"
)

// Domain contain employee behaviors
type Domain interface {
	CreateOne(employee model.Employee) error
	GetByID(id string) (*model.Employee, error)
	GetByEmail(email string) (*model.Employee, error)
	GetAll() (*model.Employees, error)
}
