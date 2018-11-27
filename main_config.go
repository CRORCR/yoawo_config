package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
	"sync"
	"time"

	"config/mode"
	"config/server"
)

func startRPCServer() {
	reward := new(server.User)
	rpc.Register(reward)
	poundage := new(server.PoundageServer)
	rpc.Register(poundage)
	buyrules := new(server.ServerBuyRules)
	rpc.Register(buyrules)
	defaulthome := new(server.ServerDHome)
	rpc.Register(defaulthome)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":7000")
	if err != nil {
		fmt.Println("错误了哦")
		os.Exit(1)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}

func taskList() {
	mode.UserInit("config/chicken_user.json")
	mode.PoundageInit("config/poundage.json")
	mode.BuyRulesInit("config/buy_rules.json")
	mode.NextHomeInit("config/next_home.json")
}

//golang 定时器，启动的时候执行一次，以后每天晚上12点执行
func startTimer(f func()) {
	go func() {
		for {
			f()
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

func main() {
	mode.GUserLock = new(sync.RWMutex)
	mode.GPounLock = new(sync.RWMutex)
	mode.GBRulLock = new(sync.RWMutex)
	mode.GHomeLock = new(sync.RWMutex)

	// 执行计划任务 5分钟执行一次
	/*
	ticker := time.NewTicker( time.Second * 300)
	go func(){
		for _ = range ticker.C{
			go mode.UserInit( "./config/chicken_user.json" )
			fmt.Println("执行计划任务")
		}
	}()
	*/
	/*
	mode.PoundageInit( "./config/poundage.json" )
	mode.GPoundage.Print()
	//mode.GPoundage.Delete( 20000 )
	var para mode.Parameter
	para.Money = 20000
	para.Poundage = 0.55
	mode.GPoundage.Add( para )
	mode.GPoundage.Set( para )
	*/
	fmt.Println("启动定时器")
	startTimer(taskList)
	fmt.Println("启动RPC")
	startRPCServer()
}
