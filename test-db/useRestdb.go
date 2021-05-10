package main

import (
	"database/sql"
	"fmt"
	"log"

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
	Hostname = "localhost"
	Port     = 5432
	Username = "mtsouk"
	Password = "pass"
	Database = "restapi"
)

func ConnectPostgres() *sql.DB {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Hostname, Port, Username, Password, Database)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Println(err)
		return nil
	}

	return db
}

func main() {
	db := restdb.ConnectPostgres()
	fmt.Println(db)
	db.Close()

	db = ConnectPostgres()
	fmt.Println(db)
	defer db.Close()

	t := restdb.User{}
	fmt.Println(t)

	rows, err := db.Query(`SELECT "username" FROM "users"`)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var username string

		err = rows.Scan(&username)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(name)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

}
