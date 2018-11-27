package server

import (
	"config/mode"
)

type PoundageServer struct {
	FileName string
}

func (this *PoundageServer) Get(inUser *mode.Poundage,
	outUser *mode.Poundage) error {
	outUser.Get()
	return nil
}

func (this *PoundageServer) SetLimit(nLimit *uint64,
	user *mode.Poundage) error {
	return user.SetLimit(*nLimit)
}

func (this *PoundageServer) Set(para *mode.Parameter,
	user *mode.Poundage) error {
	return user.Set(*para)
}

func (this *PoundageServer) Add(para *mode.Parameter,
	user *mode.Poundage) error {
	return user.Add(*para)
}

func (this *PoundageServer) Delete(nLimit *uint64,
	user *mode.Poundage) error {
	return user.Delete(*nLimit)
}
