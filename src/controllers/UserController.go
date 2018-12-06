package controllers

import (
	"base"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	base.Controller
}

var GetUserController UserController

func (this *UserController) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code":    1,
		"message": "ok",
		"data": gin.H{
			"name": "Index me",
			"age":  "99",
		},
	})
}

func (this *UserController) Create(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code":    1,
		"message": "ok",
		"data": gin.H{
			"name": "Create me",
			"age":  "99",
		},
	})
}
