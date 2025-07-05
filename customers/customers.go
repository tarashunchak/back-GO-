package CUST

import (
	DB_CONN "NeoCom/connection"
	"fmt"
)

type Customer struct {
	FName string `json:"first_name"`
	LName string `json:"last_name"`
	Phone string `json:"phone"`
	ID    int    `json:"id"`
}

func SelectAllCustomers() []Customer {
	rows, err := DB_CONN.Conn.DB.Query("SELECT first_name, last_name, phone, id FROM customers;")
	if err != nil {
		fmt.Println("SelectAllCustomers() query error")
		return nil
	}

	var customers []Customer
	for rows.Next() {
		var cust Customer
		err = rows.Scan(&cust.FName, &cust.LName, &cust.Phone, &cust.ID)
		if err != nil {
			return nil
		}

		customers = append(customers, cust)
	}
	return customers
}
