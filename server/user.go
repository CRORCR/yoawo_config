package server

import(
	"../mode"
)

type User struct {
	FileName string
}

func ( this *User )Get( nId *uint8, user *mode.User )error{
	user.Id = *nId
	user.Get()
	return nil
}

func ( this *User )Set( user *mode.User, nId *int )error{
	return user.Set()
}

