package api

import (
	"clean-arch/entity"
	"clean-arch/usercase"
)

type User struct {
	usercase.RegisterUser
}

func NewUser(userRepo entity.UserRepo, accountRepo entity.AccountRepo) (u User) {
	u.RegisterUser = usercase.NewRegisterUser(userRepo, accountRepo)
	return
}

type RegisterUserReq struct {
	Name    string
	Age     int
	Balance float64
}

type RegisterUserRes struct {
	Success bool
}

func (u User) Register(req RegisterUserReq) (res RegisterUserRes, err error) {

	return
}
