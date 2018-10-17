package main

import (
    "fmt"
    "net/rpc"
)

type User struct {
        Id              uint8
        Nuo             int
        Poundage        float64
        FreeMerchants   int
        NuoMerchants    int
}

func main(){
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

	user.Id = 2
	user.Nuo = 5000
	err = client.Call("User.Set", &user, nil )
	if err != nil {
                fmt.Println("调用失败：", err)
        }

}

