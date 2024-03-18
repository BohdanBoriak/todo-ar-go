package main

import (
	"fmt"
	"log"
	"todo-list/config/container"
	"todo-list/domain"
)

func main() {
	cont := container.New()

	u := domain.User{
		Name:     "Bohdan",
		Password: "1234",
	}

	user, err := cont.UserRepo.Save(u)
	if err != nil {
		log.Printf("Error: %s", err)
	}

	fmt.Println(user)
}
