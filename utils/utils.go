package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostData(url string, content string)  {

	var jsonStr = msg(content)
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

func msg(s string) []byte {
	jsonObj := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]string{"content": s,},
	}
	b, _ := json.Marshal(jsonObj)
	return b
}