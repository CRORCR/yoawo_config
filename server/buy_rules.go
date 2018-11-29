
package server

import(
	"config/mode"
	"fmt"
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


func (this *ServerBuyRules) GetByIndex(index int, out *mode.BuyCountAndLevel) error {
	out.Get(index)
	fmt.Println("out", out)
	return nil
}