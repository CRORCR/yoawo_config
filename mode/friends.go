package mode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// 亲密好友规则控制
type Friends struct {
	Default  int `json:"default"`   // 默认数量
	BuyCount int `json:"buy_count"` // 购买只数
	Count    int `json:"count"`     // 增加数量
}

var friends Friends

func FriendsInit(strFileName string) {
	friendLock.Lock()
	jsonFile, err := os.Open(strFileName)
	if err != nil {
		panic("打开文件错误，请查看:" + strFileName)
	}
	defer jsonFile.Close()

	jsonData, era := ioutil.ReadAll(jsonFile)
	if era != nil {
		panic("读取文件错误:" + strFileName)
	}
	json.Unmarshal(jsonData, &friends)
	fmt.Println("FILE:", friends)
	friendLock.Unlock()
}

func (this *Friends) Get() {
	*this = friends
}

var friendLock sync.Mutex

func (this *Friends) Set(def Friends) error {
	friendLock.Lock()
	defer func() { friendLock.Unlock() }()
	//可能只设置一个,防止覆盖
	friends = def
	buff, _ := json.Marshal(friends)
	err := ioutil.WriteFile("./config/friends.json", buff, 0644)
	return err
}
