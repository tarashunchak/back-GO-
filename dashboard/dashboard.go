package DASHBOARD

import (
	DB_CONN "NeoCom/connection"
	"database/sql"
)

type CustTableRowContent struct {
	Cust_ID   int    `json:"Cust. ID"`
	Full_Name string `json:"Full Name"`
	Phone     string `json:"Phone"`
	Tariff_ID int    `json:"Tariff ID"`
	Reg_date  string `json:"Reg. date"`
	Added_By  string `json:"Added by"`
	Is_Active string `json:"Is Active"`
}

type RequestsRowContent struct {
	ID    int    `json:"ID"`
	Phone string `json:"Phone"`
	Date  string `json:"Date"`
}

type RequestsBarContent struct {
	Row_Count int    `json:"field_name"`
	Date      string `json:"label"`
}

func DashboardCustomersTable(interval string) []interface{} {
	rows, err := DB_CONN.Conn.DB.Query("SELECT c.id AS \"Cust. ID\", " +
		"(c.first_name || ' ' || c.last_name) AS \"Full Name\", " +
		"c.phone AS \"Phone\", " +
		"c.tariff_id  AS \"Tariff ID\", " +
		"TO_CHAR(c.date, 'YYYY-MM-DD HH24:MI') AS \"Reg. date\", " +
		"(e.first_name || ' ' || e.last_name) AS \"Added By\", " +
		"c.is_active AS \"Is Active\" " +
		"FROM customers c " +
		"LEFT JOIN employees e ON e.id = c.added_by_id " +
		"WHERE DATE(c.date) >= DATE(CURRENT_DATE - INTERVAL '" + interval + " days') " +
		"AND c.is_visible = true " +
		"ORDER BY DATE(c.date) DESC;")

	if err != nil {
		return nil
	}

	columns, err := rows.Columns()
	var rows_list []CustTableRowContent
	for rows.Next() {
		var row_content CustTableRowContent
		rows.Scan(&row_content.Cust_ID, &row_content.Full_Name, &row_content.Phone, &row_content.Tariff_ID, &row_content.Reg_date, &row_content.Added_By, &row_content.Is_Active)

		rows_list = append(rows_list, row_content)
	}
	return []interface{}{columns, rows_list}
}

func DashboardRequestsHistory(interval string) []interface{} {
	rows, err := DB_CONN.Conn.DB.Query("SELECT r.id AS \"ID\", " +
		"c.phone AS \"Phone\", " +
		"TO_CHAR(r.date, 'YYYY-MM-DD HH24:MI') AS \"Date\" " +
		"FROM requests r " +
		"JOIN customers c ON c.id = r.cust_id " +
		"ORDER BY r.id DESC LIMIT " + interval + ";")
	if err != nil {
		return nil
	}

	columns, err := rows.Columns()
	var rows_list []RequestsRowContent
	for rows.Next() {
		var row_content RequestsRowContent
		rows.Scan(&row_content.ID, &row_content.Phone, &row_content.Date)

		rows_list = append(rows_list, row_content)
	}
	return []interface{}{columns, rows_list}
}

func DashboardRequestsSeries(interval string) []interface{} {
	var rows *sql.Rows
	var err error
	if interval == "today" {
		rows, err = DB_CONN.Conn.DB.Query("SELECT COUNT(id) AS req_count, DATE(date) AS date " +
			"FROM requests " +
			"WHERE DATE(date) = DATE(CURRENT_DATE) " +
			"GROUP BY DATE(date) ORDER BY DATE(date) DESC;")
	} else {
		rows, err = DB_CONN.Conn.DB.Query("SELECT COUNT(id) AS req_count, DATE(date) AS date " +
			"FROM requests " +
			"WHERE DATE(date) >= DATE(CURRENT_DATE - INTERVAL '" + interval + " days') " +
			"GROUP BY DATE(date) ORDER BY DATE(date) DESC;")
	}
	if err != nil {
		return nil
	}

	columns, err := rows.Columns()
	var rows_list []RequestsBarContent
	for rows.Next() {
		var row_content RequestsBarContent
		rows.Scan(&row_content.Row_Count, &row_content.Date)

		rows_list = append(rows_list, row_content)
	}
	return []interface{}{columns, rows_list}
}
