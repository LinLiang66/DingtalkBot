package handlers

import (
	"DingtalkBot/model"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func sendHelpCard(ctx context.Context,
	msgInfo *MsgInfo) {
	// 构造请求内容
	requestData := map[string]interface{}{
		"msgtype": "markdown", "markdown": map[string]interface{}{
			"title": "收到您的消息了", "text": "收到消息内容===============>>>>>" + msgInfo.QParsed,
		},
	}

	err := SendPostRequest(msgInfo.SessionWebhook, requestData)
	if err != nil {
		return
	}

}

// SendPostRequest  发起POST请求
func SendPostRequest(url string, requestBody interface{}) error {
	marshal, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("读取请求内容: %v", err)
		return err
	}

	// 构造请求对象
	req, err := http.NewRequest("POST", url, strings.NewReader(string(marshal)))
	if err != nil {
		log.Fatalf("构造请求对象: %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	// 发起请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("发起请求报错: %v", err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("读取请求结果报错: %v", err)

		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取请求结果报错: %v", err)
		return err
	}
	var res model.Result
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatalf("请求结果转换实体报错: %v", err)
		return err
	}

	return nil
}
