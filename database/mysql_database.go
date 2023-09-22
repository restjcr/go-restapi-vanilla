package database

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type MySQLDatabase struct {
	//*sql.DB
	Instance *sql.DB
}

func NewMySQLDatabase() *MySQLDatabase {
	db, err := sql.Open("mysql", "root:1234@/learning_go_apirest")

	if err != nil {
		panic(err)
	}

	return &MySQLDatabase{db}
}
