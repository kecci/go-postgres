package repository

import(
	"github.com/abyanjksatu/kecci-golang-postgresql/src/modules/employee/model"
)

//Employee Repository
type EmployeeRepository interface {
	Save(*model.Employee) error
	Update(int, *model.Employee) error
	Delete(int) error
	FindByID(int) (*model.Employee, error)
	FindAll() (model.Employees, error)
}