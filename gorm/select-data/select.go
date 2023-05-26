package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Employee struct {
	ID            uint
	Name          string
	Address       string
	Age           uint
	Birthdate     string
	Level         string
	Id_department uint
}

type Department struct {
	ID   uint
	Name string
}

func main()  {
	psqlInfo := "host=localhost port=5432 user=postgres password=1234 dbname=test-db-bismillah sslmode=disable"

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var employees []Employee

	db.Find(&employees)

	for _, employee := range employees {
		fmt.Println(employee)
	}
	var departments []Department

	db.Find(&departments)

	for _, department := range departments {
		fmt.Println(department)
	}

}