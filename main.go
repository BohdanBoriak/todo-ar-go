package main

import (
	"fmt"
	"log"
	"todo-list/config"
	"todo-list/domain"
	database "todo-list/repository"

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

	ur := database.NewUserRepository(sess)
	u := domain.User{
		Name:     "Bohdan",
		Password: "1234",
	}

	user, err := ur.Save(u)
	if err != nil {
		log.Printf("Error: %s", err)
	}

	fmt.Println(user)
}
