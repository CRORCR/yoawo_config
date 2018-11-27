package mode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type ChickenReward struct {
	Number int `json:"number"` // 奖励的鸡的编号
	Count  int `json:"count"`  // 奖励的数量
}

// 注册新用户奖励机制
type RegNewUser struct {
	Cash        float64         `json:"cash"`         // 现金
	Trust       float64         `json:"trust"`        // 鍩分
	Feeding     int             `json:"feeding"`      // 食物
	Water       int             `json:"water"`        // 水
	Band        int             `json:"band"`         // 鍩豆
	DHome       int             `json:"default_home"` // 下家数量限制
	ChickenList []ChickenReward `json:"chicken"`      // 鸡的奖励列表
}

var GNewUser RegNewUser
var GNewUserLock *sync.RWMutex

/*
 * 描述：把相应配置文件中的数据刷新到内存中，本方法在
 *
 *	计划任务中执行。
 *
 **************************************************************/
func RegUserInit(strFileName string) {
	GNewUserLock.RLock()
	jsonFile, err := os.Open(strFileName)
	if err != nil {
		panic("打开文件错误，请查看:" + strFileName)
	}
	defer jsonFile.Close()

	jsonData, era := ioutil.ReadAll(jsonFile)
	if era != nil {
		panic("读取文件错误:" + strFileName)
	}
	json.Unmarshal(jsonData, &GNewUser)
	fmt.Println("FILE:", GNewUser)
	GNewUserLock.RUnlock()
}

/*
 * 描述：根据nType数值不同返回不同的数
 *
 **************************************************************/
//func ( this *GNewUser )Get(){
//	*this = GNewUser
//}
