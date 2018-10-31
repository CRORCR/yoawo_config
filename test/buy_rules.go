
package main

import (
    "fmt"
    "net/rpc"
)

type ServerBuyRules struct {
        FileName string
}

type BuyRules struct {
        Name            string  `json:"name"`           // 规则名称
        Number          int     `json:"number"`         // 规则编号
        Enable          bool    `json:"enable"`         // 是否启用
        V45             int     `json:"v45"`            // 45天伐值
        V60             int     `json:"v60"`            // 60天伐值
        V75             int     `json:"v75"`             // 75天伐值
        V90             int     `json:"v90"`             // 90天伐值
}

type ModeBuyRules struct {
        Rules []BuyRules        `json:"rules"`
}



func main(){
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	var user ModeBuyRules
	err = client.Call("ServerBuyRules.Get", &user, &user )
	if err != nil {
		fmt.Println("调用失败：", err)
	}
	fmt.Println("调用结果：", user)

}

