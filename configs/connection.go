package configs

import (
	"database/sql"
	"fmt"
	"log"

	// driver mysql for go
	_ "github.com/go-sql-driver/mysql"
)

// Dbmap variable de conectividad
var Dbmap = connectionConfig()

const (
	user     = ""
	password = ""
	dbname   = "apirest_users"
)

func connectionConfig() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8", user, password, dbname))
	if err != nil {
		log.Fatal("Error al conectar a la db: ", err.Error())
	}
	return db
}
