
package server

import(
	"fmt"
	"config/mode"
)

type ServerBuyRules struct {
	FileName string
}

func ( this *ServerBuyRules )Get( inRules *mode.ModeBuyRules, outRules *mode.ModeBuyRules )error{
	outRules.Get()
	fmt.Println("aaaaaaaaaaa", outRules)
	return nil
}

func ( this *ServerBuyRules )Set( inRules *mode.BuyRules, outRules *mode.ModeBuyRules )error{
	outRules.Set(inRules)
	return nil
}
