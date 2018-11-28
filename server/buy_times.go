package server

import (
	"config/mode"
)

type ServerBuyTimes struct {
	FileName string
}

func (this *ServerBuyTimes) Get(index , outRules *mode.BuyTimes) error {
	outRules.Get()
	//fmt.Println("index", index)
	//fmt.Println("out", outRules)
	return nil
}

func (this *ServerBuyTimes) Set(inRules *mode.BuyTimes, outRules *mode.BuyTimes) error {
	outRules.Set(inRules)
	return nil
}