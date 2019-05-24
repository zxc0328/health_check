package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

var timeHeuristic = time.Duration(8 * time.Second)

var TIMEOUT = timeHeuristic

func main() {
	if err := MakeRequest(); err != nil {
		SendAlert("外包项目挂了哟")
	}
}

func SendAlert(text string) {
	log.Println("发送警报")
	message := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": text,
		},
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	_, err = http.Post("https://oapi.dingtalk.com/robot/send?access_token=0fc384d57235fdb1cc6dfa83408d8154507d0699f3649589dd5a7898012ee690", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Println(err)
	}
}

// account.ccnu.edu.cn 模拟登录，用于验证账号密码是否可以正常登录
func MakeRequest() error {

	// 初始化 http client
	client := http.Client{
		Timeout: TIMEOUT,
	}

	// 初始化 http request
	request, err := http.NewRequest("GET", "https://upyingtou.com/sd/health", nil)
	if err != nil {
		log.Println(err)
		return err
	}

	// 发起请求
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("Status error")
	}

	return nil
}
