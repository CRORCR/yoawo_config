package server

import (
	"config/mode"
)

type ServerHighLimit struct {
	FileName string
}

func (this *ServerHighLimit) Get(sdefault *mode.HighObject) error {
	sdefault.Get()
	return nil
}

func (this *ServerHighLimit) Set(nType, sdefault *mode.HighObject) error {
	sdefault.Set(nType)
	return nil
}
