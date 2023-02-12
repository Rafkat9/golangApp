package adapters

import "github.com/gin-gonic/gin"

type Handler interface {
	Register(ctx *gin.Context)
}
