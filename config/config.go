package config

/*

key: gitlab 的 http 地址
value: 企业微信机器人的地址

*/
var GitLabHookMap = map[string]string{
	"https://gitlab.luojilab.com/fanwenqiang/DD-Project.git": "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=e591ec2b-4849-4eca-8e77-4b93dd907ec9",
	"https://gitlab.luojilab.com/igc-ios/IGCLive.git": "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=e591ec2b-4849-4eca-8e77-4b93dd907ec9",
	"https://gitlab.luojilab.com/mazhensai/test.git": "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=f1d7cfa5-f4dd-4ee5-bfc0-b58a6e5f5409",
}