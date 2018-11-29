package server

import (
	"config/mode"
	"fmt"
)

type ServerBuyTimes struct {
	FileName string
}

func (this *ServerBuyTimes) Get(index int, a *mode.Int) error {
	a.Get(index)
	fmt.Println("out", *a)
	return nil
}

func (this *ServerBuyTimes)Set(inRules *mode.BuyTimes, outRules *mode.BuyTimes) error {
	outRules.Set(inRules)
	return nil
}