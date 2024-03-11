package main

import (
	"fmt"
	"log"
	"todo-list/config"

	"github.com/upper/db/v4/adapter/postgresql"
)

func main() {
	config := config.GetConfig()

	var settings = postgresql.ConnectionURL{
		Database: config.DbName,
		Host:     config.DbHost,
		User:     config.DbUser,
		Password: config.DbPassword,
	}

	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatalf("Couldn't establish db connection")
	}

	fmt.Println(sess)
}
