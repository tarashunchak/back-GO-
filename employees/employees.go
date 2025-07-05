package EMPL

import (
	DB_CONN "NeoCom/connection"
	"fmt"
)

type Employee struct {
	FName string `json:"first_name"`
	LName string `json:"last_name"`
	ID    int    `json:"id"`
}

func SelectAllEmployees() []Employee {
	rows, err := DB_CONN.Conn.DB.Query("SELECT first_name, last_name, id FROM employees;")

	if err != nil {
		fmt.Println("SelectAllEmployees query error")
		return nil
	}

	var employees []Employee
	for rows.Next() {
		var empl Employee

		err = rows.Scan(&empl.FName, &empl.LName, &empl.ID)

		if err != nil {
			return nil
		}
		employees = append(employees, empl)
	}

	return employees
}

func SelectEmployeesBy(str string) []Employee {
	rows, err := DB_CONN.Conn.DB.Query("SELECT first_name, last_name, id FROM employees "+
		"WHERE id = $1;", str)

	if err != nil {
		return nil
	}

	var employees []Employee
	for rows.Next() {
		var empl Employee

		err = rows.Scan(&empl.FName, &empl.LName, &empl.ID)

		if err != nil {
			return nil
		}
		employees = append(employees, empl)
	}

	return employees
}
