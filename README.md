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

修改 `config.json` 中的 `gitlabHooks`

```json
{
      "gitlabURL": "your/gitlab/repo/path",
      "weChatRobotURL": "your/weChatWork/Bot/URL",
      "push": true,
      "merge": true,
      "tag": true
}
```
字段含义

1. push 是否处理 push
2. merge 是否处理 merge
3. tag 是否处理 tag

#### life

修改 `config.json` 中的 `life`，在数组中添加如下 json 对象

```json
 {
      "weChatRobotURL": "your/weChatWork/Bot/URL",
      "drinkWater": true,
      "offDuty": true,
      "orderMeal": true,
      "pee": true
 }
```

字段含义

1. drinkWater 是否提醒喝水
2. offDuty 是否提醒下班
3. orderMeal 是否提醒订饭
4. pee 是否提醒嘘嘘

#### 部署
1. 编译并部署到服务器上
2. 在 `gitlab webhook` 上，填入服务器接口地址 `you.server.domain:9091/gitlabHook`

欢迎提交 `issue` 或者 `Pull Request`