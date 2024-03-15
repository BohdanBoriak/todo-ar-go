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

func (r UserRepository) FindById(id uint64) (domain.User, error) {
	var u user
	err := r.coll.Find("id = ?", id).One(&u)
	if err != nil {
		return domain.User{}, err
	}
	usr := r.mapModelToDomain(u)
	return usr, nil
}

func (r UserRepository) Update(u domain.User) (domain.User, error) {
	usr := r.mapDomainToModel(u)
	err := r.coll.Find("id = ?", u.Id).Update(&usr)
	if err != nil {
		return domain.User{}, err
	}
	updatedUser := r.mapModelToDomain(usr)
	return updatedUser, nil
}

func (r UserRepository) Delete(id uint64) error {
	err := r.coll.Find("id = ?", id).Delete()
	return err
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
