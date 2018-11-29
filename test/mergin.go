package main

import (
    "fmt"
    "net/rpc"
)

type Margin struct {
	Nuo   int `json:"nuo"`   // nuo余额
	Money int `json:"money"` // 现金余额
}
func main(){
	//getMargin()
	setMargin()
}
func setMargin() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	defer func() {client.Close()}()
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	var margin = Margin{0, 0}
	err = client.Call("ServerMargin.Set", &margin, &margin)

	if err != nil {
		fmt.Println("调用失败：", err)
	}
	fmt.Println("调用结果：", margin)
}


func getMargin() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	var margin Margin
	err = client.Call("ServerMargin.Get", &margin, &margin)
	if err != nil {
		fmt.Println("调用失败：", err)
	}
	fmt.Println("调用结果：", margin)
}

