
package server

import(
	"config/mode"
)

type ServerBuyRules struct {
	FileName string
}

func ( this *ServerBuyRules )Get( inRules *mode.BuyRules, outRules *mode.BuyRules )error{
	outRules.Get()
	return nil
}

func ( this *ServerBuyRules )Set( inRules, outRules *mode.BuyRules)error{
	outRules.Set(inRules)
	return nil
}