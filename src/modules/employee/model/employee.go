package model

import(
	"time"
)

type GenderType string

const (
    M GenderType = "M"
    F GenderType = "F"
)

//Employee Struct
type Employee struct {
	ID int
	Name string
	Birthdate time.Time
	Gender GenderType
	Married bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Employees type Employee List
type Employees []Employee

//NewEmployee Employee's Constructor
func NewEmployee() *Employee{
	return &Employee{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}