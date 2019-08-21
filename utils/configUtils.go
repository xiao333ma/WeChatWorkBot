package utils

import (
	"encoding/json"
	"io/ioutil"
)

type configModel struct {
	GitlabHooks []GitlabHooksModel `json:"gitlabHooks"`
	Life []LifeModel `json:"life"`
}

type GitlabHooksModel struct {
	GitlabURL      string `json:"gitlabURL"`
	WeChatRobotURL string `json:"weChatRobotURL"`
	Push bool `json:"push"`
	Merge bool `json:"merge"`
	Tag bool `json:"tag"`
}

type LifeModel struct {
	WeChatRobotURL string `json:"weChatRobotUrl"`
	DrinkWater bool `json:"drinkWater"`
	OffDuty bool `json:"offDuty"`
	OrderMeal bool `json:"orderMeal"`
	Pee bool `json:"pee"`
}

func GetGitLabWeChatRobotURL(gitURL string) GitlabHooksModel {
	config, err := readConfigJSON()
	if err != nil {
		return GitlabHooksModel{}
	}
	for _, value := range config.GitlabHooks {
		if value.GitlabURL == gitURL {
			return value
		}
	}
	return GitlabHooksModel{}
}

func GetLifeWeChatRobotURL() []LifeModel {
	config, err := readConfigJSON()
	if err != nil {
		return []LifeModel{}
	}
	return config.Life
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