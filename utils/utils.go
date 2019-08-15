package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostData(url string, content string, atAll bool)  {
	jsonStr := markdownMsg(content)
	post(url, jsonStr)
	if atAll {
		atMessage := textMessage("请大佬们处理👆👆👆")
		post(url, atMessage)
	}
}

func post(url string, jsonStr []byte)  {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func markdownMsg(s string) []byte {
	jsonObj := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]string{"content": s,},
	}
	b, _ := json.Marshal(jsonObj)
	return b
}

func textMessage(s string) []byte  {
	jsonObj := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": s,
			"mentioned_mobile_list": []string{"@all"},
		},
	}
	b, _ := json.Marshal(jsonObj)
	return b
}