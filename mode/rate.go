package mode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

/**
 * @desc    行业利率开关
 * @author Ipencil
 * @create 2018/11/30
 */

type RateList struct{
	RateOf []*Rate `json:"rate"`
}

type Rate struct {
	TopName  string `json:"top_name"`
	SuperName string `json:"super_name"`
	Gas float64 `json:"gas"`
}

var (
	rateList RateList
)

func RateListInit(strFileName string) {
	rateLock.Lock()
	jsonFile, err := os.Open(strFileName)
	if err != nil {
		panic("打开文件错误，请查看:" + strFileName)
	}
	defer func() {
		jsonFile.Close()
		rateLock.Unlock()
	}()

	jsonData, era := ioutil.ReadAll(jsonFile)
	if era != nil {
		panic("读取文件错误:" + strFileName)
	}
	json.Unmarshal(jsonData, &rateList)
	fmt.Println("rateList:", rateList)
}

func (this *RateList) Get() {
	*this = rateList
}

var rateLock sync.Mutex

func (this *RateList) Set(rateList *RateList) error {
	fmt.Printf("set mode:%+v\n", rateList)
	rateLock.Lock()
	//覆盖文件
	fileName:="config/rate_list.json"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC, 0644)
	defer func() {
		file.Close()
		rateLock.Unlock()
		}()
	if err != nil {
		return err
	}
	buff, _ := json.Marshal(rateList)
	_, err = file.Write(buff)
	if err != nil {
		return err
	}
	fmt.Println("write", len(buff))

	return err
}
