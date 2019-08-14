package router

import (
	"WeChatWorkRobot/handler"
	"fmt"
	"github.com/gin-gonic/gin"
)

var postRouter = map[string]handler.HttpBusinessHandler{
	"/gitlabHook": handler.NewGitlabHookHandler(),
}

func Router()  {
	router := gin.Default()

	for path, h := range postRouter {
		router.POST(path, h.HandleBusiness)
	}

	err := router.Run(":9091")
	if err != nil {
		fmt.Println("发生错误❌", err)
	}
}
