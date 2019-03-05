package controller

import(
	"fmt"
	"encoding/json"
	"log"
	// "math/rand"
	"net/http"
	// "strconv"
	"github.com/gorilla/mux"

	// "github.com/abyanjksatu/kecci-golang-postgresql/src/modules/employee/model"
	"github.com/abyanjksatu/kecci-golang-postgresql/src/modules/employee/repository"
	"github.com/abyanjksatu/kecci-golang-postgresql/config"
)

// Get all employees
func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//init Database
	db, err := config.GetPostgresDB()

	if err != nil {
		fmt.Println(err)
	}

	employeeRepositoryPostgres := repository.NewProfileRepositoryPostgres(db)

	employees, err := employeeRepositoryPostgres.FindAll()

	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(employees)
}

func HandleRequest() error {
	// Init Router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/employee", getEmployees).Methods("GET")
	// r.HandleFunc("/employee/{id}", getEmployee).Methods("GET")
	// r.HandleFunc("/employee", createEmployee).Methods("POST")
	// r.HandleFunc("/employee/{id}", updateEmployee).Methods("PUT")
	// r.HandleFunc("/employee/{id}", deleteEmployee).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))

	return nil
}