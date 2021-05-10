package main

import (
	"fmt"

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

func main() {
	db := restdb.ConnectPostgres()
	fmt.Println(db)

	t := restdb.User{}
	fmt.Println(t)
}
