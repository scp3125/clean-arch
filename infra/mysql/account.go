package mysql

import "clean-arch/entity"

type account struct {
}

func NewAccount() (a account) {

	return
}

func (u account) FindOneByID(id int64) (account entity.Account, err error) {
	return
}

func (u account) Create(account entity.Account) (id int64, err error) {
	return
}

func (u account) Update(account entity.Account) (id int64, err error) {
	return
}

func (u account) Delete(id int64) (err error) {
	return
}
