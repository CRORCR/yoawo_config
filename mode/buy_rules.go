package mode

import (
	"config/lib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
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
	Limit  int    `json:"limit"`  //购鸡总额控制
}

var (
	buyRules  BuyRules
	GBRulLock *sync.RWMutex
	buyCountAndLevel BuyCountAndLevel
)

type BuyCountAndLevel struct {
	Level int `json:"level"`
	Num   int `json:"num"`
}

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
	json.Unmarshal(jsonData, &buyRules)
	fmt.Println("FILE:", &buyRules)
	GBRulLock.RUnlock()
}

/*
 * 描述：修改相应的数据
 *
 **************************************************************
func ( this *IdentityReward )Set()error{

	// STEP 1 加锁
	GBRulLock.Lock()

	// STEP 2 设置修改数据
	var user User
	var k int
	for k,_ = range GUser {
		if GUser[k].Id == this.Id {
			user = GUser[k]
			GUser[k] = *this
			break
		}
	}

	// STEP 3 写入到文件
	buff, _ := json.Marshal( GUser )
	err := ioutil.WriteFile( "./config/chicken_user.json", buff, 0644 )
	GUser[k] = user

	// STEP 4 解锁
	GBRulLock.Unlock()
	return err
}
*/

/*
 * desc:  返回指定额度等级
 * @create: 2018/11/27
 */
func (a *BuyCountAndLevel) Get(index int) {
	//传入的参数处于哪个级别,返回可购数量
	fmt.Println("index=10000?",index)
	*a = getLevelAndCount(index)
	fmt.Println("a==?",a)
	return
}

/*
 * desc:  "v45":150,"v60":600,"v75":7000,"v90":90000
 * @create: 2018/11/28
 */
func getLevelAndCount(index int) (a BuyCountAndLevel) {
	buy :=BuyCountAndLevel{}
	fmt.Println("getLevelAndCount=10000?",index)
	if index<150{
		buy.Level=0
		buy.Num=0
	}else if index>=lib.BuyLevelAndCountArray[0][1] && index<lib.BuyLevelAndCountArray[1][1]{//150-600
		buy.Level=lib.BuyLevelAndCountArray[0][0]
		v:=index/lib.BuyLevelAndCountArray[0][1]
		float,_ := strconv.ParseFloat(strconv.Itoa(v),64)
		buy.Num=int(math.Floor(float))
	}else if index>=lib.BuyLevelAndCountArray[1][1] && index<lib.BuyLevelAndCountArray[2][1]{//600-7000
		buy.Level=lib.BuyLevelAndCountArray[1][0]
		v:=index/lib.BuyLevelAndCountArray[1][1]
		float,_ := strconv.ParseFloat(strconv.Itoa(v),64)
		buy.Num=int(math.Floor(float))
	}else if index>=lib.BuyLevelAndCountArray[2][1] && index<lib.BuyLevelAndCountArray[3][1]{
		buy.Level=lib.BuyLevelAndCountArray[2][0]
		v:=index/lib.BuyLevelAndCountArray[2][1]
		float,_ := strconv.ParseFloat(strconv.Itoa(v),64)
		buy.Num=int(math.Floor(float))
	}else{
		buy.Level=lib.BuyLevelAndCountArray[3][0]
		v:=index/lib.BuyLevelAndCountArray[3][1]
		float,_ := strconv.ParseFloat(strconv.Itoa(v),64)
		buy.Num=int(math.Floor(float))
	}
	return buy
}

/*
 * 描述：返回所有的节点数据
 *
 **************************************************************/
func (this *BuyRules) Get() {
	*this = buyRules
	fmt.Println("this", this)
	fmt.Println("grules", buyRules)
}

func (this *BuyRules) Set(rules *BuyRules) error {
	fmt.Printf("set mode:%+v\n", rules)
	GBRulLock.Lock()
	if buyRules.Number == rules.Number {
		buyRules = *rules
	}
	buff, _ := json.Marshal(buyRules)
	err := ioutil.WriteFile("./config/buy_rules.json", buff, 0644)
	GBRulLock.Unlock()
	return err
}

/*
 * 描述：添加本节点数据
 *
 **************************************************************
func ( this *IdentityIdentityReward )Add()error{
	var fage bool = false
	if 0 != this.Id {
		for _,v := range GUser {
			if v.Id == this.Id {
				fage = true
				break
			}
		}
		if fage {
			GUser = append( GUser, *this )
			return this.Set()
		}
	}
	return nil
}*/
