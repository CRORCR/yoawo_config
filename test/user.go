package main

import (
	"fmt"
	"net/rpc"
)

type User struct {
	Id            uint8   `json:"user"`
	Nuo           int     `json:"nuo"`
	Poundage      float64 `json:"poundage"`
	FreeMerchants int     `json:"free_merchants"`
	NuoMerchants  int     `json:"nuo_merchants"`
}

func main() {
	getUser()
	//setUser()
}

func setUser() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")
	var user User
	user.Id = 2  //id相同会覆盖
	user.Nuo = 1001
	err = client.Call("User.Set", &user, nil)
	//其他参数不填会置空
	if err != nil {
		fmt.Println("调用失败：", err)
	}
}

func getUser() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	var user User
	var a uint8
	a = 2
	err = client.Call("User.Get", &a, &user )
	if err != nil {
		fmt.Println("调用失败：", err)
	}
	fmt.Println("调用结果：", user)
}