
package main

import (
    "fmt"
    "net/rpc"
)

type DefaultHmoe struct {
	DefaultHome	int	`json:"default_home"`		// 新用户默认给多少下家的数量
	ChickenHome	int	`json:"chicken_home"`		// 购买一只鸡给多少下家的数量
}

func main(){
	//getLimit()
	setLimit()
}
func setLimit() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	var user = DefaultHmoe{1, 1}
	err = client.Call("ServerDHome.Set", &user, &user)
	if err != nil {
		fmt.Println("调用失败：", err)
	}
	fmt.Println("调用结果：", user)
}
func getLimit() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	var user DefaultHmoe
	err = client.Call("ServerDHome.Get", &user, &user)
	if err != nil {
		fmt.Println("调用失败：", err)
	}
	fmt.Println("调用结果：", user)
}

