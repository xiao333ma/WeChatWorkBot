package handler

import (
	"WeChatWorkRobot/gitlabHook/model"
	"WeChatWorkRobot/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type gitlabHookHandler struct {
	HttpBusinessHandler
	weChatBotURL string
}

func NewGitlabHookHandler() *gitlabHookHandler {
	return &gitlabHookHandler{}
}

type eventHandler func (model model.GitLabHookModel)

var eventHandlerMap = map[string]eventHandler{
	"push": pushHandler,
	"merge_request": mergeRequestHandler,
	"tag_push": tagPushHandler,
}

func (handler *gitlabHookHandler)HandleBusiness (c *gin.Context)  {
	var m = model.GitLabHookModel{}
	err := c.BindJSON(&m)

	if err != nil {
		fmt.Println("转换错误")
	}
	preProcess(&m)
	f, exist := eventHandlerMap[m.Name()]

	if exist {
		f(m)
	}
}

func pushHandler(hookModel model.GitLabHookModel)  {
	m  := utils.GetGitLabWeChatRobotURL(hookModel.Project.GitHttpUrl)
	if !m.Push {
		return
	}
	robot := m.WeChatRobotURL
	msgContent := utils.Title(3, "有人 push 代码啦 👏 ") + utils.Newline()
	msgContent += utils.GreenString(hookModel.UserName)
	msgContent += utils.WhiteSpace() + "在" + utils.WhiteSpace()
	msgContent += getLastCommitTime(hookModel) + utils.WhiteSpace() + "提交了代码" + utils.Newline()
	msgContent += "信息如下" + utils.Newline()
	msgContent += "👉 仓库: " + utils.Link(hookModel.Project.Name, hookModel.Project.GitHttpUrl) + utils.Newline()
	msgContent += "👉 分支: " + utils.GreenString(hookModel.Ref) + utils.Newline()
	msgContent += "👉 信息: " + utils.Newline()
	msgContent += commitMessage(hookModel)
	msgContent += "你可以点击" + utils.WhiteSpace()
	msgContent += utils.Link("这里", getLastCommitURL(hookModel))
	msgContent += utils.WhiteSpace() + "查看"
	utils.PostMarkdownData(robot, msgContent)
}

func mergeRequestHandler(hookModel model.GitLabHookModel) {

	m  := utils.GetGitLabWeChatRobotURL(hookModel.Project.GitHttpUrl)
	if !m.Merge {
		return
	}
	robot := m.WeChatRobotURL
	title := "有人"
	option := model.MergeAction[hookModel.ObjectAttributes.Action]
	title += option + "了一个 Merge Request 👏"

	msgContent := utils.Title(4, title) + utils.Newline()
	msgContent += utils.GreenString(hookModel.User.Name)
	msgContent += utils.WhiteSpace() + "在" + utils.WhiteSpace()
	msgContent += formatMergeRequestTime(hookModel.ObjectAttributes.CreatedAt) + utils.WhiteSpace() + option + "了一个 Merge Request" + utils.Newline()
	msgContent += "信息如下" + utils.Newline()
	msgContent += "👉 仓库: " + utils.Link(hookModel.Project.Name, hookModel.Project.GitHttpUrl) + utils.Newline()
	msgContent += "👉 分支: " + utils.GreenString(hookModel.ObjectAttributes.SourceBranch + " --> " + hookModel.ObjectAttributes.TargetBranch) + utils.Newline()
	msgContent += "👉 信息: " + hookModel.ObjectAttributes.Title + utils.Newline()
	msgContent += "你可以点击" + utils.WhiteSpace()
	msgContent += utils.Link("这里", hookModel.ObjectAttributes.Url)
	msgContent += utils.WhiteSpace() + "查看 批准 评论"
	utils.PostMarkdownData(robot, msgContent)
	utils.PostTextData(robot, "请大佬们处理👆👆👆", true)
}

func tagPushHandler(hookModel model.GitLabHookModel)  {
	m  := utils.GetGitLabWeChatRobotURL(hookModel.Project.GitHttpUrl)
	if !m.Merge {
		return
	}
	robot := m.WeChatRobotURL

	title := "有人"
	option := ""
	if len(hookModel.Commits) > 0 {
		option = "创建"
	} else {
		option = "关闭"
	}
	title += option + "了一个 Tag 👏"

	msgContent := utils.Title(3, title) + utils.Newline()
	msgContent += utils.GreenString(hookModel.UserName)
	msgContent += utils.WhiteSpace() + option + "了 Tag" + utils.Newline()
	msgContent += "信息如下" + utils.Newline()
	msgContent += "👉 仓库: " + utils.Link(hookModel.Project.Name, hookModel.Project.GitHttpUrl) + utils.Newline()
	msgContent += "👉 Tag: " + utils.GreenString(hookModel.Ref) + utils.Newline()
	msgContent += "👉 信息: " + hookModel.Message + utils.Newline()
	utils.PostMarkdownData(robot, msgContent)
}

func preProcess(hookModel *model.GitLabHookModel) {
	hookModel.Ref = getCommitBranch(hookModel.Ref)
}


func getCommitBranch(s string) string  {
	arr := strings.Split(s, "/")
	return arr[len(arr) - 1]
}

func getLastCommitTime(hookModel model.GitLabHookModel) string  {
	return formatPushTime(getLastCommit(hookModel).Timestamp)
}

func getLastCommitURL(hookModel model.GitLabHookModel) string  {

	return getLastCommit(hookModel).Url
}

func getLastCommit(hookModel model.GitLabHookModel) model.GitLabHookModelCommits  {
	return hookModel.Commits[len(hookModel.Commits) - 1]
}

func formatPushTime(stringTime string) string  {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02T15:04:05Z", stringTime, loc)
	s := theTime.Add(time.Hour * 8).Format("2006-01-02 15:04:05")
	return s
}

func formatMergeRequestTime(stringTime string) string {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05 UTC", stringTime, loc)
	s := theTime.Add(time.Hour * 8).Format("2006-01-02 15:04:05")
	return s
}

func commitMessage(hookModel model.GitLabHookModel) string  {
	line := ""
	for i := len(hookModel.Commits) - 1; i >= 0; i-- {
		s := hookModel.Commits[i].Message
		s = strings.Replace(s, "#", "", -1)
		arr := strings.Split(s, "\n")
		for _, value := range arr {
			line += utils.Ref(value) + utils.Newline()
		}
	}
	return line + utils.Newline()
}
