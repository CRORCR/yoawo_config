
package server

import(
	"config/mode"
)

type ServerMargin struct {
	FileName string
}

func ( this *ServerMargin )Get( nType, sdefault *mode.Margin )error{
	sdefault.Get()
	return nil
}

func ( this *ServerMargin )Set(nType,sdefault *mode.Margin )error{
	sdefault.Set(*nType)
	return nil
}
