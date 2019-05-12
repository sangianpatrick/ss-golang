package model

// Employee struct contains employee's property
type Employee struct {
	ID        string `bson:"employeeID" json:"employeeID"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
	Email     string `bson:"email" json:"email"`
	Phone     string `bson:"phone" json:"phone"`
	Address   string `bson:"address" json:"address"`
	Position  string `bson:"position" json:"position"`
}

// Employees contains list of employee
type Employees []Employee
