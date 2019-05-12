/**
 * oop_basic1
 */

package main

import "fmt"

// User is a struct that contains user's property
type User struct {
	ID    string
	Name  string
	Email string
	Age   uint
}

// NewUser is a user's constructor
func NewUser(id string, name string, email string, age uint) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
		Age:   age,
	}
}

// GetName returns user's name
func (u *User) GetName() string {
	return u.Name
}

func main() {
	fmt.Printf("Golang OOP Basic\n")

	user := NewUser("USER0001", "Patrick", "patrick@email", 17)
	fmt.Printf("User's name is %s\n", user.Name)
	fmt.Printf("Detail of User: %v\n", user)
}
