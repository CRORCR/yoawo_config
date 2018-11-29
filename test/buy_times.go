package main

import (
	"fmt"
	"net/rpc"
)

type BuyTimes struct {
	V500  int `json:"v500"`
	V1K   int `json:"v1000"`
	V2K   int `json:"v2000"`
	V5K   int `json:"v5000"`
	V20K  int `json:"v20000"`
	V50K  int `json:"v50000"`
	V100K int `json:"v100000"`
	V150K int `json:"v150000"`
}

type Int int

type ServerBuyTimes struct {
	FileName string
}

func main() {
	getTimes()
	//setTimes()
}
func setTimes() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	var rules BuyTimes
	var user BuyTimes
	rules.V500 = 500000
	rules.V1K = 10
	rules.V2K = 0
	rules.V5K = 0
	rules.V20K = 10
	rules.V50K = 8
	rules.V100K = 5
	rules.V150K = 3
	//number必须对应上,否则就是增加了
	err = client.Call("ServerBuyTimes.Set", &rules, &user)
	if err != nil {
		fmt.Println("调用失败：", err)
	}
}

func getTimes() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	var index = 100
	var result Int
	err = client.Call("ServerBuyTimes.Get", &index, &result)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Println("调用结果:", result)
}
