package server

import (
	"config/mode"
)

type ServerRateList struct {
	FileName string
}

func (this *ServerRateList) Get(nType,sdefault *mode.RateList) error {
	sdefault.Get()
	return nil
}

func (this *ServerRateList) Set(nType, sdefault *mode.RateList) error {
	sdefault.Set(nType)
	return nil
}
