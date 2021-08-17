package service

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"spicychicken.top/chaos-b/dao"
	"spicychicken.top/chaos-b/model"
	"spicychicken.top/chaos-b/pkgs/mysql"
)

const ()

// CreateChaosTest install instance with custom mycnf file
func CreateChaosTest(serverAddr, cmd string) (response []byte, err error) {
	tx := mysql.GormDB.Begin()
	// 在开始测试首先需要创建相应的测试记录
	dao.CreateTest(tx, &model.Test{})
	// 将测试的结果更新到建立的测试记录中

	return
}

func sendTestRequest(serverAddr, cmd string) (response []byte, err error) {
	resp, err := http.Get(fmt.Sprintf("%s/chaosblade?cmd=%s", serverAddr, cmd))
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if response, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
}
