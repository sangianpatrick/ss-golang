package repository

import (
	"github.com/sangianpatrick/ss-golang/oop_basic3/modules/employee"
	"github.com/sangianpatrick/ss-golang/oop_basic3/modules/employee/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// EmployeeMongoRepository contains employee's entity with mongodb
type EmployeeMongoRepository struct {
	sess *mgo.Session
}

// NewEmployeeMongoRepository acts like constructor
func NewEmployeeMongoRepository(mgoSess *mgo.Session) (employee.MongoRepository, error) {
	err := mgoSess.Ping()
	if err != nil {
		return nil, err
	}
	return &EmployeeMongoRepository{
		sess: mgoSess,
	}, nil
}

// InsertOne will add a new employee.
func (mr *EmployeeMongoRepository) InsertOne(employee model.Employee) error {
	return mr.sess.DB("gotest").C("employee").Insert(employee)
}

// FindByID returns employee property
func (mr *EmployeeMongoRepository) FindByID(id string) (*model.Employee, error) {
	var employee model.Employee
	var query = bson.M{
		"employeeID": id,
	}
	err := mr.sess.DB("gotest").C("employee").Find(query).One(&employee)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

// FindByEmail returns employee property
func (mr *EmployeeMongoRepository) FindByEmail(email string) (*model.Employee, error) {
	var employee model.Employee
	var query = bson.M{
		"email": email,
	}
	err := mr.sess.DB("gotest").C("employee").Find(query).One(&employee)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

// FindAll return all employee
func (mr *EmployeeMongoRepository) FindAll() (*model.Employees, error) {
	var employees model.Employees
	var query = bson.M{}
	err := mr.sess.DB("gotest").C("employee").Find(query).All(&employees)
	if err != nil {
		return nil, err
	}
	return &employees, nil
}
