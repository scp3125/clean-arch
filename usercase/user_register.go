package usercase

import "clean-arch/entity"

// 注册用户
type RegisterUser struct {
	userRepo    entity.UserRepo
	accountRepo entity.AccountRepo
}

func NewRegisterUser(userRepo entity.UserRepo, accountRepo entity.AccountRepo) (r RegisterUser) {
	r.userRepo = userRepo
	r.accountRepo = accountRepo
	return
}

func (r RegisterUser) Register(userName string, userAge int, balance float64) (err error) {
	userId, err := r.userRepo.Create(entity.User{Name: userName, Age: userAge})
	if err != nil {
		return
	}
	_, err = r.accountRepo.Create(entity.Account{UserID: userId, Balance: balance})
	if err != nil {
		return
	}
	return
}
