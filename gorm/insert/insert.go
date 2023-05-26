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

	var departments = []Department{
		{Name: "TI Department"},
		{Name: "Economy Department"},
		{Name: "Science Department"},
		{Name: "Political Department"},
		{Name: "Social Department"},
	}
	db.Create(&departments)

	var employees = []Employee{
		{
			Name:          "Rinnai",
			Address:       "Ciledug",
			Age:           25,
			Birthdate:     "2001-11-21",
			Level:         "Secretary",
			Id_department: 1,
		},
		{
			Name:          "Rajin",
			Address:       "Cikokol",
			Age:           12,
			Birthdate:     "1995-11-21",
			Level:         "CEO",
			Id_department: 2,
		},
		{
			Name:          "Nirna",
			Address:       "Ciamis",
			Age:           23,
			Birthdate:     "1995-11-21",
			Level:         "CTO",
			Id_department: 3,
		},
		{
			Name:          "Raniah",
			Address:       "Bandung",
			Age:           29,
			Birthdate:     "1995-11-21",
			Level:         "HRD",
			Id_department: 4,
		},
		{
			Name:          "Rumiati",
			Address:       "Surabaya",
			Age:           18,
			Birthdate:     "1995-11-21",
			Level:         "Secretary",
			Id_department: 5,
		},
	}

	db.Create(&employees)

	fmt.Println("Insertion Success")

}