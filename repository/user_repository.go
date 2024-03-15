package database

import (
	"todo-list/domain"

	"github.com/upper/db/v4"
)

const usersTableName = "users"

type user struct {
	Id       uint64 `db:"id,omitempty"`
	Name     string `db:"name"`
	Password string `db:"password"`
}

type UserRepository struct {
	coll db.Collection
	sess db.Session
}

func NewUserRepository(sess db.Session) UserRepository {
	return UserRepository{
		coll: sess.Collection(usersTableName),
		sess: sess,
	}
}

func (r UserRepository) Save(u domain.User) (domain.User, error) {
	usr := r.mapDomainToModel(u)
	err := r.coll.InsertReturning(&usr)
	if err != nil {
		return domain.User{}, err
	}
	newUser := r.mapModelToDomain(usr)
	return newUser, err
}

func (r UserRepository) mapDomainToModel(u domain.User) user {
	return user{
		Id:       u.Id,
		Name:     u.Name,
		Password: u.Password,
	}
}

func (r UserRepository) mapModelToDomain(u user) domain.User {
	return domain.User{
		Id:       u.Id,
		Name:     u.Name,
		Password: u.Password,
	}
}
