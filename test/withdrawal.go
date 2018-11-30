package main

import (
	"encoding/json"
	"fmt"
	"net/rpc"
)

type ServerHighLimit struct {
	FileName string
}
type HighObject struct {
	Highest  int     `json:"highest"`
	DrawList []*Draw `json:"draw_list"`
}

type Draw struct {
	Low  int `json:"low"`
	High int `json:"high"`
	Gas  int `json:"gas"`
}

/**
 * @desc    用户提现开关
 * @author Ipencil
 * @create 2018/11/30
 */
func main() {
	setStruct()
	//getHighLimit()
	setHighLimit()
}
func setHighLimit() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	err = client.Call("ServerHighLimit.Set", &highObject, nil)

	if err != nil {
		fmt.Println("调用失败：", err)
	}
	fmt.Println("调用结果：", highObject)
}

func getHighLimit() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	var highObject HighObject
	err = client.Call("ServerHighLimit.Get", &highObject, &highObject)
	if err != nil {
		fmt.Println("调用失败：", err)
	}
	fmt.Printf("调用结果:%+v", highObject)
}

var highObject = &HighObject{}

func setStruct() {
	dr := make([]*Draw, 0)
	dr = append(dr, &Draw{10, 100, 1})
	dr = append(dr, &Draw{20, 200, 2})
	dr = append(dr, &Draw{30, 300, 2})
	dr = append(dr, &Draw{40, 400, 4})
	dr = append(dr, &Draw{50, 500, 5})
	highObject.Highest = 100000000
	highObject.DrawList = dr
	bytes, _ := json.Marshal(highObject)
	fmt.Println(string(bytes))
}
