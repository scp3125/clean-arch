package entity

type Account struct {
	ID      int64
	UserID  int64
	Balance float64
	Margin  float64
}

type AccountRepo interface {
	FindOneByID(id int64) (account Account, err error)
	Create(account Account) (id int64, err error)
	Update(account Account) (id int64, err error)
	Delete(id int64) (err error)
}

// 能否进行交易
func (a Account) CanDeal() bool {
	return a.Margin >= 1000
}
