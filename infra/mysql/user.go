package mysql

import "clean-arch/entity"

type user struct {
}

func NewUser() (u user) {

	return
}

func (u user) FindOneByID(id int64) (user entity.User, err error) {
	return
}

func (u user) Create(user entity.User) (id int64, err error) {
	return
}

func (u user) Update(user entity.User) (id int64, err error) {
	return
}

func (u user) Delete(id int64) (err error) {
	return
}
