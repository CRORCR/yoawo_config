package mode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// 推广奖励
type IdentityReward struct {
	Id            uint8   `json:"user"`
	Nuo           int     `json:"nuo"`
	Poundage      float64 `json:"poundage"`
	FreeMerchants int     `json:"free_merchants"`
	NuoMerchants  int     `json:"nuo_merchants"`
}

var GUser []IdentityReward
var GUserLock *sync.RWMutex

/*
 * 描述：	把相应配置文件中的数据刷新到内存中，本方法在
 *
 *	计划任务中执行。
 *
 **************************************************************/
func UserInit(strFileName string) {
	GUserLock.RLock()
	jsonFile, err := os.Open(strFileName)
	if err != nil {
		panic("打开文件错误，请查看:" + strFileName)
	}
	defer jsonFile.Close()

	jsonData, era := ioutil.ReadAll(jsonFile)
	if era != nil {
		panic("读取文件错误:" + strFileName)
	}
	json.Unmarshal(jsonData, &GUser)
	GUserLock.RUnlock()
}

/*
 * 描述：打印结构
 *
 **************************************************************/
func (this *IdentityReward) Print() {
	fmt.Println("用户身份编号:", this.Id)
	fmt.Println("鍩分    奖励:", this.Nuo)
	fmt.Println("提现    奖励:", this.Poundage)
	fmt.Println("普通商家奖励:", this.FreeMerchants)
	fmt.Println("鍩分商家奖励:", this.NuoMerchants)
}

/*
 * 描述：修改相应的数据
 *
 **************************************************************/
func (this *IdentityReward) Set() error {

	// STEP 1 加锁
	GUserLock.Lock()

	// STEP 2 设置修改数据
	var user IdentityReward
	var k int
	for k, _ = range GUser {
		if GUser[k].Id == this.Id {
			user = GUser[k]
			GUser[k] = *this
			break
		}
	}

	// STEP 3 写入到文件
	buff, _ := json.Marshal(GUser)
	err := ioutil.WriteFile("./config/chicken_user.json", buff, 0644)
	GUser[k] = user

	// STEP 4 解锁
	GUserLock.Unlock()
	return err
}

/*
 * 描述：获取相应的节点数据
 *
 **************************************************************/
func (this *IdentityReward) Get() {
	for k, _ := range GUser {
		if GUser[k].Id == this.Id {
			this.Nuo = GUser[k].Nuo
			this.Poundage = GUser[k].Poundage
			this.FreeMerchants = GUser[k].FreeMerchants
			this.NuoMerchants = GUser[k].NuoMerchants
		}
	}
}

/*
 * 描述：添加本节点数据
 *
 **************************************************************/
func (this *IdentityReward) Add() error {
	var fage bool = false
	if 0 != this.Id {
		for _, v := range GUser {
			if v.Id == this.Id {
				fage = true
				break
			}
		}
		if fage {
			GUser = append(GUser, *this)
			return this.Set()
		}
	}
	return nil
}
