package database

import "github.com/upper/db/v4"

type user struct {
	Id       uint64 `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
}

type UserRepository struct {
	coll db.Collection
	sess db.Session
}
