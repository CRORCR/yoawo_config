package main

import (
    "fmt"
    "net/rpc"
)

// 亲密好友规则控制
type Friends struct {
	Default  int `json:"default"`   // 默认数量
	BuyCount int `json:"buy_count"` // 购买只数
	Count    int `json:"count"`     // 增加数量
}

func main(){
	//getFriends()
	setFriends()
}
func setFriends() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	defer func() {client.Close()}()
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	var friends = Friends{1, 0,0}
	err = client.Call("ServeFriends.Set", &friends, &friends)

	if err != nil {
		fmt.Println("调用失败：", err)
	}
	fmt.Println("调用结果：", friends)
}


func getFriends() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	var friends Friends
	err = client.Call("ServeFriends.Get", &friends, &friends)
	if err != nil {
		fmt.Println("调用失败：", err)
	}
	fmt.Println("调用结果：", friends)
}

