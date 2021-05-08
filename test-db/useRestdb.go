package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/mactsouk/restdb"
)

type User struct {
	ID        int
	Username  string
	Password  string
	LastLogin int64
	Admin     int
	Active    int
}

// PostgreSQL Connection details
var (
	Hostname = "postgres"
	Port     = 5432
	Username = "mtsouk"
	Password = "pass"
	Database = "restapi"
)

func main() {
	var db *sql.DB
	db = restdb.ConnectPostgres()
}
