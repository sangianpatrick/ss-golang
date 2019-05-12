package employee

import (
	"github.com/sangianpatrick/ss-golang/oop_basic3/modules/employee/model"
)

// MongoRepository contains behavior of employee mongodb entity.
type MongoRepository interface {
	InsertOne(employee model.Employee) error
	FindByID(id string) (*model.Employee, error)
	FindByEmail(email string) (*model.Employee, error)
	FindAll() (*model.Employees, error)
}
