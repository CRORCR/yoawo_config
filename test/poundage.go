package main

import (
    "fmt"
    "net/rpc"
)


type PoundageServer struct {
	FileName string
}

type Parameter struct {
	Money	 uint64	 `json:"money"`
	Poundage float64 `json:"poundage"`
}

type Poundage struct {
	MonthLimit	uint64		`json:"MonthLimit"`
	NodeList	[]Parameter     `json:"NoteList"`
}


func main(){
	client, err := rpc.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	var usera Poundage
	var userb Poundage
	err = client.Call("PoundageServer.Get", &usera, &userb )
	if err != nil {
		fmt.Println("调用失败：", err)
	}
	fmt.Println("调用结果：", userb)

}

