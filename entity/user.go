package entity

import (
	"crypto/sha256"
	"encoding/hex"
)

type User struct {
	ID   int64
	Name string
	Age  int
}

type UserRepo interface {
	FindOneByID(id int64) (user User, err error)
	Create(user User) (id int64, err error)
	Update(user User) (id int64, err error)
	Delete(id int64) (err error)
}

// 获取名字哈希值
func (u User) GetNameHash() (hash string) {
	h := sha256.New()
	h.Write([]byte(u.Name))
	sum := h.Sum(nil)
	hash = hex.EncodeToString(sum)
	return
}
