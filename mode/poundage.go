package mode

import (
	"os"
	"fmt"
	"sync"
	"io/ioutil"
	"encoding/json"
)

type Parameter struct {
	Money	  uint64  `json:"money"`
	Poundage  float64 `json:"poundage"`
}

type Poundage struct {
	MonthLimit	uint64		`json:"MonthLimit"`
	NodeList	[]Parameter     `json:"NoteList"`
}

var GPoundage	Poundage
var GPounLock	*sync.RWMutex

/*
 * 描述:
 *		把相应配置文件中的数据刷新到内存中，本方法在
 *	计划任务中执行。
 *
 *************************************************************/
func PoundageInit( strFileName string ) {
	GPounLock.RLock()
	jsonFile, err := os.Open(strFileName)
	if err != nil {
                panic("打开文件错误，请查看:" + strFileName)
        }
        defer jsonFile.Close()

        jsonData, era := ioutil.ReadAll(jsonFile)
        if era != nil {
                panic("读取文件错误:" + strFileName)
        }
        json.Unmarshal( jsonData, &GPoundage )
	GPounLock.RUnlock()
}

/*
 * 描述：打印结构
 *
 *************************************************************/
func ( this *Poundage )Print(){
	fmt.Println("用户提现最高限额:", this.MonthLimit )
	fmt.Println( "*************************************" )
	for _,v := range this.NodeList {
		fmt.Println( "提现限额", v.Money )
		fmt.Println( "手 续 费", v.Poundage )
		fmt.Println( "*****************************" )
	}
}

/*
 * 描述: 写到文件中
 *
 *************************************************************/
func ( this *Poundage )save() error {
	// STEP 1 加锁
	GPounLock.Lock()

	// STEP 2 写入数据
	buff, _ := json.Marshal( this )
	err := ioutil.WriteFile( "./config/poundage.json",
				 buff,
				 0644 )

	// STEP 3 解锁
	GPounLock.Unlock()
	return err
}

func ( this *Poundage )Get(){
	*this = GPoundage
}

/*
 * 描述：设置用户每月提现最高限额
 *
 *************************************************************/
func ( this *Poundage )SetLimit( nLimit uint64 )error{
	this.Get()
	this.MonthLimit = nLimit
	return this.save()
}

/*
 * 描述：
 *
 *************************************************************/
func ( this *Poundage )Set( sPara Parameter )error{
	this.Get()
	for _, v := range this.NodeList {
		if v.Money == sPara.Money{
			this.Delete( sPara.Money )
			this.Add( sPara )
		}
	}
	return this.save()
}

/*
 * 描述：添加提现节点
 *
 *************************************************************/
func ( this *Poundage )Add( sNode Parameter )error{
	this.Get()
	var fage bool = true
	for _, v := range this.NodeList{
		if v.Money == sNode.Money {
			fage = false
		}
	}
	if fage {
		this.NodeList = append( this.NodeList,sNode )
		return this.save()
	}
	return nil
}

/*
 * 描述：删除提现节点
 *
 *************************************************************/
func ( this *Poundage )Delete( nLimit uint64 ) error {
	this.Get()
	for k, v := range this.NodeList {
		if v.Money == nLimit {
			this.NodeList = append(
				this.NodeList[:k],
				this.NodeList[k+1:]...)
		}
	}
	return this.save()
}

