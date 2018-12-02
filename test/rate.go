package main

import (
	"encoding/json"
	"fmt"
	"net/rpc"
)
/**
 * @desc    行业利率
 * @author Ipencil
 * @create 2018/11/30
 */

type RateList struct{
	RateOf []*Rate `json:"rate"`
}

type Rate struct {
	TopName  string `json:"top_name"`
	SuperName string `json:"super_name"`
	Gas float64 `json:"gas"`
}

func main() {
	//getRateList()
	setRate()
	setRateList()
}

func setRateList() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	err = client.Call("ServerRateList.Set", &rateList, &rateList)

	if err != nil {
		fmt.Println("调用失败：", err)
	}
	fmt.Println("调用结果：", rateList)
}

func getRateList() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	var rateList RateList
	err = client.Call("ServerRateList.Get", &rateList, &rateList)
	if err != nil {
		fmt.Println("调用失败：", err)
	}
	fmt.Printf("调用结果:%+v", rateList)
}

var rates = make([]*Rate,0)
var rateList = RateList{}
func setRate(){
	rates=append(rates,&Rate{"餐饮/食品","餐饮",float64(0.3)})
	rates=append(rates,&Rate{"餐饮/食品","食品",float64(0.2)})
	rates=append(rates,&Rate{"线下零售","超市",float64(0.3)})
	rates=append(rates,&Rate{"线下零售","便利店",float64(0.4)})
	rates=append(rates,&Rate{"教育/医疗","私立院校",float64(0.5)})
	rateList.RateOf=rates
	bytes, _ := json.Marshal(rateList)
	fmt.Println(string(bytes))
}