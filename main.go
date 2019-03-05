package main

import (
	"fmt"
	// "github.com/abyanjksatu/kecci-golang-postgresql/config"
	"github.com/abyanjksatu/kecci-golang-postgresql/src/controller"
)

func main(){
	fmt.Println("Golang Postgresql")

	//init Database
	// db, err := config.GetPostgresDB()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	err := controller.HandleRequest();

	if err != nil {
		fmt.Println(err)
	}
}