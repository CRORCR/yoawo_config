
package server

import(
	"config/mode"
)

type ServerDHome struct {
	FileName string
}

func ( this *ServerDHome )Get( nType, sdefault *mode.DefaultHmoe )error{
	sdefault.Get()
	return nil
}

func ( this *ServerDHome )Set(nType,sdefault *mode.DefaultHmoe )error{
	sdefault.Set(*nType)
	return nil
}
