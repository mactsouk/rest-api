package main

import (
	"fmt"
	"log"
	"time"

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
	user := restdb.User{ID: 0, Username: "mtsouk", Password: "admin", LastLogin: time.Now().Unix(), Admin: 1, Active: 1}
	if restdb.InsertUser(user) {
		fmt.Println("User inserted successfully.")
	} else {
		fmt.Println("Insert failed!")
	}

	mtsoukUser := restdb.FindUserUsername(user.Username)
	fmt.Println("mtsouk: ", mtsoukUser)

	if restdb.DeleteUser(mtsoukUser.ID) {
		fmt.Println("User Deleted.")
	} else {
		fmt.Println("User not Deleted.")
	}

	mtsoukUser = restdb.FindUserUsername(user.Username)
	fmt.Println("mtsouk: ", mtsoukUser)

	if restdb.DeleteUser(mtsoukUser.ID) {
		fmt.Println("User Deleted.")
	} else {
		fmt.Println("User not Deleted.")
	}

	if restdb.DeleteUser(mtsoukUser.ID) {
		fmt.Println("User Deleted.")
	} else {
		fmt.Println("User not Deleted.")
	}
}
