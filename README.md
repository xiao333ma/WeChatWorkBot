# 企业微信机器人

### Feature

Gitlab

- [x] push
- [x] merge request
- [x] tag

Life

- [x] 提醒喝水
- [x] 提醒订饭
- [x] 提醒下班
- [x] 嘘嘘 🐔

### Usage

#### gitlabHooks

1. 修改 `config.json` 中的 `gitlabHooks`
```json
{
      "gitlabURL": "your/gitlab/repo/path",
      "weChatRobotURL": "your/weChatWork/Bot/URL",
      "push": true or false to handle push event,
      "merge": true or false to handle merge event,
      "tag": true or false to handle tag event
}
```

#### life

1. 修改 `config.json` 中的 `life`，在数组中添加如下 json 对象

```json
 {
      "weChatRobotURL": "your/weChatWork/Bot/URL",
      "drinkWater": true or false to alert drink water,
      "offDuty": true or false to alert off duty,
      "orderMeal": true or false to alert order meal,
      "pee": true or false to alert pee
 }
```

#### 部署
1. 编译并部署到服务器上
2. 在 `gitlab webhook` 上，填入服务器接口地址 `you.server.domain:9091/gitlabHook`

欢迎提交 `issue` 或者 `Pull Request`