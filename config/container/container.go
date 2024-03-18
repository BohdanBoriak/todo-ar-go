package container

import (
	"log"
	"todo-list/config"
	database "todo-list/repository"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

type Container struct {
	UserRepo database.UserRepository
}

func New() Container {
	c := config.GetConfig()
	sess := getDbConnection(c)

	ur := database.NewUserRepository(sess)
	return Container{
		UserRepo: ur,
	}
}

func getDbConnection(c config.Config) db.Session {
	var settings = postgresql.ConnectionURL{
		Database: c.DbName,
		Host:     c.DbHost,
		User:     c.DbUser,
		Password: c.DbPassword,
	}
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatalf("Couldn't establish db connection: %s", err)
	}
	return sess
}
