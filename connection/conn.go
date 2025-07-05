package DB_CONN

import (
	"database/sql"
	"fmt"
)

type Connection struct {
	DB *sql.DB
}

func (con *Connection) ConnectToDB() {
	connStr := "postgresql://neondb_owner:npg_qILNuP6Diz1Z@ep-divine-sun-a83zg48v-pooler.eastus2.azure.neon.tech/neondb?sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Cannot connect to DB")
		return
	}
	con.DB = db
}

var Conn Connection
