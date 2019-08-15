# 企业微信机器人

### Feature

Gitlab

- [x] push
- [x] merge request
- [x] tag

Life

- [ ] 提醒喝水
- [ ] 提醒订饭

### Usage

1. 修改 `config.go`, `key` 为 `gitlab webhook` 地址，`value` 为企业微信`机器人地址` 
1. 把代码部署到服务器上
2. 在 `gitlab webhook` 上，填入服务器接口地址 `you.server.domain:9091/gitlabHook`



想加什么功能，请提交 `issue` 或者 `Merge Request`