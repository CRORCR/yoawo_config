package mode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// 购买规制
type DefaultHmoe struct {
	DefaultHome int `json:"default_home"` // 新用户默认给多少下家的数量
	ChickenHome int `json:"chicken_home"` // 购买一只鸡给多少下家的数量
}

var GNextHome DefaultHmoe
var GHomeLock *sync.RWMutex

/*
 * 描述：	把相应配置文件中的数据刷新到内存中，本方法在
 *
 *	计划任务中执行。
 *
 **************************************************************/
func NextHomeInit(strFileName string) {
	GHomeLock.RLock()
	jsonFile, err := os.Open(strFileName)
	if err != nil {
		panic("打开文件错误，请查看:" + strFileName)
	}
	defer jsonFile.Close()

	jsonData, era := ioutil.ReadAll(jsonFile)
	if era != nil {
		panic("读取文件错误:" + strFileName)
	}
	json.Unmarshal(jsonData, &GNextHome)
	fmt.Println("FILE:", GNextHome)
	GHomeLock.RUnlock()
}

/*
 * 描述：根据nType数值不同返回不同的数
 *
 **************************************************************/
func (this *DefaultHmoe) Get() {
	*this = GNextHome
}
func (this *DefaultHmoe) Set(def DefaultHmoe) error {
	GUserLock.Lock()
	GNextHome = def
	// STEP 3 写入到文件
	buff, _ := json.Marshal(GNextHome)
	err := ioutil.WriteFile("./config/next_home.json", buff, 0644)

	// STEP 4 解锁
	GUserLock.Unlock()
	return err
}
