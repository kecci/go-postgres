package repository

import(
	"database/sql"
	"github.com/abyanjksatu/kecci-golang-postgresql/src/modules/employee/model"
)

type employeeRepositoryPostgres struct {
	db *sql.DB
}

func NewProfileRepositoryPostgres(db *sql.DB) *employeeRepositoryPostgres{
	return &employeeRepositoryPostgres{db}
}

func (r *employeeRepositoryPostgres) Save(employee *model.Employee) error {
	query :=  `INSERT INTO "employee" ("id", "name", "birthdate", "gender", "married", "createdAt", "updatedAt")
	VALUES($1, $2, $3, $4, $5, $6, $7) `

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(employee.ID, employee.Name, employee.Birthdate, employee.Gender, employee.Married, employee.CreatedAt, employee.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *employeeRepositoryPostgres) Update(id int, employee *model.Employee) error {
	query := `UPDATE "employee" SET name"=$1, "birthdate"=$2, "gender"=$3, "married"=$4, "updatedAt"=$5 WHERE "id"=$6`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(employee.Name, employee.Birthdate, employee.Gender, employee.Married, employee.UpdatedAt, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *employeeRepositoryPostgres) Delete(id int) error {
	query := `DELETE FROM "employee" WHERE "id" = $1`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func (r *employeeRepositoryPostgres) FindById(id int) (*model.Employee, error) {
	query := `SELECT * FROM "employee" WHERE "id" = $1`

	var employee model.Employee

	statement, err := r.db.Prepare(query)	

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(&employee.ID, &employee.Name, &employee.Birthdate, &employee.Gender, &employee.Married, &employee.CreatedAt, &employee.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &employee, nil
}


func(r *employeeRepositoryPostgres) FindAll() (model.Employees, error) {
	query := `SELECT * FROM public.employee`

	var employees model.Employees

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next(){
		var employee model.Employee

		err = rows.Scan(&employee.ID, &employee.Name, &employee.Birthdate, &employee.Gender, &employee.Married, &employee.CreatedAt, &employee.UpdatedAt)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil

}











