package main

import (
	"fmt"
	"net/rpc"
)

type Parameter struct {
	Money    uint64  `json:"money"`
	Poundage float64 `json:"poundage"`
}

type Poundage struct {
	MonthLimit uint64      `json:"month_limit"`
	NodeList   []Parameter `json:"note_list"`
}

func main() {
	//getPoundage()
	setPoundage()
}
func setPoundage() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	var para = Parameter{5000, 100}
	var userb Poundage
	err = client.Call("PoundageServer.Set", &para, &userb)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Println("调用结果:", userb)
}

func getPoundage() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")
	var usera Poundage
	var userb Poundage
	err = client.Call("PoundageServer.Get", &usera, &userb)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Println("调用结果:", userb)
}
