package mode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

/**
 * @desc    用户体现开关
 * @author Ipencil
 * @create 2018/11/30
 */

type HighObject struct{
	Highest int `json:"highest"`
	DrawList []*Draw `json:"draw_list"`
}

type Draw struct {
	Low  int `json:"low"`
	High int `json:"high"`
	Gas int `json:"gas"`
}

var (
	highObject HighObject
)

func HighObjectInit(strFileName string) {
	highLock.Lock()
	jsonFile, err := os.Open(strFileName)
	if err != nil {
		panic("打开文件错误，请查看:" + strFileName)
	}
	defer func() {
		jsonFile.Close()
		highLock.Unlock()
	}()

	jsonData, era := ioutil.ReadAll(jsonFile)
	if era != nil {
		panic("读取文件错误:" + strFileName)
	}
	json.Unmarshal(jsonData, &highObject)
	fmt.Println("highObject:", highObject)
}

func (this *HighObject) Get() {
	fmt.Println("into high get")
	fmt.Println("object:",highObject)
	*this = highObject
}

var highLock sync.Mutex

func (this *HighObject) Set(highObject *HighObject) error {
	fmt.Printf("set mode:%+v\n", highObject)
	highLock.Lock()
	//覆盖文件
	fileName:="config/withdrawal.json"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC, 0644)
	defer func() {
		file.Close()
		highLock.Unlock()
		}()
	if err != nil {
		return err
	}
	buff, _ := json.Marshal(highObject)
	_, err = file.Write(buff)
	if err != nil {
		return err
	}
	fmt.Println("write", len(buff))

	return err
}
