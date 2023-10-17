package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	host     = ""
	port     = 5432
	user     = ""
	password = ""
	dbname   = "test"
)

func main() {
	fmt.Println("nomodo.io postgresql connection test")

	var err error

	fmt.Print("host: ")
	fmt.Scanf("%s", &host)

	fmt.Print("port: ")
	if _, err := fmt.Scanf("%d", &port); err != nil {
		log.Fatal(err)
	}

	fmt.Print("user: ")
	fmt.Scanf("%s: ", &user)

	fmt.Print("password: ")
	fmt.Scanf("%s ", &password)

	psqlInfo := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/?sslmode=disable",
		user, password, host, port,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully connected!")
}
