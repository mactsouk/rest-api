package main

import (
	"fmt"
	"log"
	"time"

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

func main() {
	db := restdb.ConnectPostgres()
	fmt.Println(db)
	defer db.Close()

	err := db.Ping()
	if err != nil {
		fmt.Println("Ping:", err)
		return
	}

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

		fmt.Println(username)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("Populating PostgreSQL")
	user := restdb.User{0, "mtsouk", "admin", time.Now().Unix(), 1, 0}
	if restdb.InsertUser(user) {
		fmt.Println("User inserted successfully.")
	} else {
		fmt.Println("Insert failed!")
	}

}
