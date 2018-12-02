package main

import (
	"fmt"
	"net/rpc"
)

type ServerBuyRules struct {
	FileName string
}

type BuyRules struct {
	Name   string `json:"name"`   // 规则名称
	Number int    `json:"number"` // 规则编号
	Enable bool   `json:"enable"` // 是否启用
	V45    int    `json:"v45"`    // 45天伐值
	V60    int    `json:"v60"`    // 60天伐值
	V75    int    `json:"v75"`    // 75天伐值
	V90    int    `json:"v90"`    // 90天伐值
	Limit  int    `json:"limit"`   //购鸡总额控制
}

type ModeBuyRules struct {
	Rules []BuyRules `json:"rules"`
}

type BuyCountAndLevel struct {
	Level int `json:"level"`
	Num   int `json:"num"`
}

func main() {
	//getBuy()
	//setbBuy()

	getBuyByCount()
}
func setbBuy() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	var rules BuyRules
	rules.Number = 0
	//number必须对应上,否则就是增加了
	rules.Enable = true
	rules.V45 = 12345
	err = client.Call("ServerBuyRules.Set", &rules, &rules)
	if err != nil {
		fmt.Println("调用失败：", err)
	}
}

func getBuy() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	var buyRules BuyRules
	err = client.Call("ServerBuyRules.Get", &buyRules, &buyRules)
	if err != nil {
		fmt.Println("调用失败：", err)
	}
	fmt.Println("调用结果:", buyRules)
}

//"v45":500,"v60":700,"v75":7000,"v90":9000,
func getBuyByCount() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")

	var index = 6000
	var result BuyCountAndLevel
	err = client.Call("ServerBuyRules.GetByIndex", &index, &result)

	if err != nil {
		fmt.Println("调用失败：", err)
	}
	fmt.Println("调用结果:", result)
}