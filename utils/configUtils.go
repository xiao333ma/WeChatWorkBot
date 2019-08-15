package utils

import (
	"encoding/json"
	"io/ioutil"
)

type configModel struct {
	GitlabHooks []GitlabHooksModel `json:"gitlabHooks"`
}

type GitlabHooksModel struct {
	GitlabURL      string `json:"gitlabURL"`
	WeChatRobotURL string `json:"weChatRobotURL"`
}


func GetWeChatRobotURL(gitURL string) string {
	config, err := readConfigJSON()
	if err != nil {
		return ""
	}
	for _, value := range config.GitlabHooks {
		if value.GitlabURL == gitURL {
			return value.WeChatRobotURL
		}
	}
	return ""
}

func readConfigJSON() (configModel, error) {
	config := configModel{}

	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return config, err
	}
	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &config)
	return config, err
}