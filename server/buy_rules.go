
package server

import(
	"fmt"
	"../mode"
)

type ServerBuyRules struct {
	FileName string
}

func ( this *ServerBuyRules )Get( inRules *mode.ModeBuyRules, outRules *mode.ModeBuyRules )error{
	outRules.Get()
	fmt.Println("aaaaaaaaaaa", outRules)
	return nil
}

