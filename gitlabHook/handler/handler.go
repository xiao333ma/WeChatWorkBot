package handler

import "github.com/gin-gonic/gin"

type HttpBusinessHandler interface {
	HandleBusiness(c *gin.Context)
}