package mode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// 购买规制
type Margin struct {
	Nuo   int `json:"nuo"`   // nuo余额
	Money int `json:"money"` // 现金余额
}

var GMarginHome Margin

/*
 * 描述：	把相应配置文件中的数据刷新到内存中，本方法在
 *
 *	计划任务中执行。
 *
 **************************************************************/
func MarginInit(strFileName string) {
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
	json.Unmarshal(jsonData, &GMarginHome)
	fmt.Println("FILE:", GMarginHome)
	GHomeLock.RUnlock()
}

/*
 * 描述：根据nType数值不同返回不同的数
 *
 **************************************************************/
func (this *Margin) Get() {
	*this = GMarginHome
}

var lock sync.Mutex

func (this *Margin) Set(def Margin) error {
	lock.Lock()
	defer func() { lock.Unlock() }()
	//可能只设置一个,防止覆盖
	if def.Money > 0 {
		GMarginHome.Money = def.Money
	}
	if def.Nuo > 0 {
		GMarginHome.Nuo = def.Nuo
	}
	buff, _ := json.Marshal(GMarginHome)
	err := ioutil.WriteFile("./config/margin.json", buff, 0644)
	return err
}
