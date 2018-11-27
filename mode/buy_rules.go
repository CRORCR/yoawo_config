package mode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// 购买规制
type BuyRules struct {
	Name   string `json:"name"`   // 规则名称
	Number int    `json:"number"` // 规则编号
	Enable bool   `json:"enable"` // 是否启用
	V45    int    `json:"v45"`    // 45天伐值
	V60    int    `json:"v60"`    // 60天伐值
	V75    int    `json:"v75"`    // 75天伐值
	V90    int    `json:"v90"`    // 90天伐值
}

type ModeBuyRules struct {
	Rules []BuyRules `json:"rules"`
}

var GBuyRules ModeBuyRules
var GBRulLock *sync.RWMutex

/*
 * 描述：	把相应配置文件中的数据刷新到内存中，本方法在
 *
 *	计划任务中执行。
 *
 **************************************************************/
func BuyRulesInit(strFileName string) {
	GBRulLock.RLock()
	jsonFile, err := os.Open(strFileName)
	if err != nil {
		panic("打开文件错误，请查看:" + strFileName)
	}
	defer jsonFile.Close()

	jsonData, era := ioutil.ReadAll(jsonFile)
	if era != nil {
		panic("读取文件错误:" + strFileName)
	}
	json.Unmarshal(jsonData, &GBuyRules)
	fmt.Println("FILE:", &GBuyRules)
	GBRulLock.RUnlock()
}

/*
 * 描述：修改相应的数据
 *
 */
func (this *ModeBuyRules) Set(rules *BuyRules) error {
	fmt.Printf("set mode:%+v\n",rules)
	// STEP 1 加锁
	GBRulLock.Lock()
	// STEP 2 设置修改数据
	for k, _ := range GBuyRules.Rules {
		if GBuyRules.Rules[k].Number == rules.Number {
			GBuyRules.Rules[k]=*rules
		}
	}
	// STEP 3 写入到文件
	buff, _ := json.Marshal(GBuyRules)
	err := ioutil.WriteFile("./config/buy_rules.json", buff, 0644)
	// STEP 4 解锁
	GBRulLock.Unlock()
	return err
}

/*
 * desc:  返回所有的节点数据
 * @create: 2018/11/27
 */
func (this *ModeBuyRules) Get() {
	*this = GBuyRules
	fmt.Println("this", this)
	fmt.Println("grules", GBuyRules)
}
