package server

import (
	"config/mode"
)

type User struct {
	FileName string
}

func (this *User) Get(nId *uint8, user *mode.IdentityReward) error {
	user.Id = *nId
	user.Get()
	return nil
}

func (this *User) Set(user *mode.IdentityReward, nId *int) error {
	return user.Set()
}
