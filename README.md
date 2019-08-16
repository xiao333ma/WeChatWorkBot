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

### Usage

#### gitlabHooks

1. 修改 `config.json` 中的 `gitlabHooks`
`gitlabURL` 为 仓库地址
`weChatRobotURL` 为 企业微信机器人地址

#### life

1. 修改 `config.json` 中的 `life`，在数组中添加`企业微信机器人地址`

#### 部署
1. 把代码部署到服务器上
2. 在 `gitlab webhook` 上，填入服务器接口地址 `you.server.domain:9091/gitlabHook`

想加什么功能，请提交 `issue` 或者 `Pull Request`