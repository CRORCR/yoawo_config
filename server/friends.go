package server

import (
	"config/mode"
)

type ServeFriends struct {
	FileName string
}

func (this *ServeFriends) Get(nType, sdefault *mode.Friends) error {
	sdefault.Get()
	return nil
}

func (this *ServeFriends) Set(nType, sdefault *mode.Friends) error {
	sdefault.Set(*nType)
	return nil
}
