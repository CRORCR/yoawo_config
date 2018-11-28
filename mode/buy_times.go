package mode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

/*
500以下，收益倍数最高5倍；
500-1000，收益倍数最高10倍；
1000-2000，收益倍数，不限制；
2000-5000，收益倍数，不限制；
5000-2万，收益倍数最高10倍；
2万－5万，收益倍数最高8倍；
5万-10万，收益倍数最高为5倍；
10万－15万，收益倍数最高为3倍；
15万以上，收益倍数最高为2倍；
 */
// 购买规制
type BuyTimes struct {
	V500  int `json:"v500"`
	V1K   int `json:"v1000"`
	V2K   int `json:"v2000"`
	V5K   int `json:"v5000"`
	V20K  int `json:"v20000"`
	V50K  int `json:"v50000"`
	V100K int `json:"v100000"`
	V150K int `json:"v150000"`
}

type Int int

var (
	buyTimes  BuyTimes
	GBBuyLock *sync.RWMutex
	fileName  string
)

/*
 * 描述：	把相应配置文件中的数据刷新到内存中，本方法在
 *
 *	计划任务中执行。
 *
 **************************************************************/
func BuyTimesInit(strFileName string) {
	GBRulLock.RLock()
	fileName = strFileName
	jsonFile, err := os.Open(strFileName)
	if err != nil {
		panic("打开文件错误，请查看:" + strFileName)
	}
	defer jsonFile.Close()

	jsonData, era := ioutil.ReadAll(jsonFile)
	if era != nil {
		panic("读取文件错误:" + strFileName)
	}
	err = json.Unmarshal(jsonData, &buyTimes)
	fmt.Println("FILE:", &buyTimes)
	GBRulLock.RUnlock()
}

/*
 * 描述：修改相应的数据
 *
 */
func (this *BuyTimes) Set(rules *BuyTimes) error {
	fmt.Printf("set mode:%+v\n", rules)
	GBRulLock.Lock()
	//覆盖文件
	fmt.Println("filename", fileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC, 0644)
	defer func() { file.Close() }()
	if err != nil {
		return err
	}
	buff, _ := json.Marshal(rules)
	_, err = file.Write(buff)
	if err != nil {
		return err
	}
	fmt.Println("write", len(buff))
	GBRulLock.Unlock()
	return err
}

/*
 * desc:  返回所有的节点数据
 * @create: 2018/11/27
 */
func (this *BuyTimes) Get() {
	// STEP 2 设置修改数据
	//fmt.Println("index==?",index)
	*this = buyTimes
	fmt.Println("this:", this)
	return
}

func getLevel(index int) int {
	if index > 150000 { //大于15w
		return buyTimes.V150K
	} else if index > 100000 && index <= 150000 { //大于10w-15w
		return buyTimes.V100K
	} else if index > 50000 && index <= 100000 { //大于5w-10万
		return buyTimes.V5K
	} else if index > 20000 && index <= 500000 {
		return buyTimes.V20K
	} else if index > 5000 && index <= 200000 {
		return buyTimes.V5K
	} else if index > 2000 && index <= 5000 {
		return buyTimes.V2K
	} else if index > 1000 && index <= 2000 {
		return buyTimes.V2K
	} else if index > 500 && index <= 1000 {
		return buyTimes.V1K
	} else {
		return buyTimes.V500
	}
}

/*
500以下，收益倍数最高5倍；
500-1000，收益倍数最高10倍；
1000-2000，收益倍数，不限制；
2000-5000，收益倍数，不限制；
5000-2万，收益倍数最高10倍；
2万－5万，收益倍数最高8倍；
5万-10万，收益倍数最高为5倍；
10万－15万，收益倍数最高为3倍；
15万以上，收益倍数最高为2倍；
 */
