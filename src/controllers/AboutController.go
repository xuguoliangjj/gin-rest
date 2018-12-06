package controllers

import (
	"base"
	"github.com/gin-gonic/gin"
)

type AboutController struct {
	base.Controller
}

var GetAboutController AboutController

func (this *AboutController) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code":    1,
		"message": "ok",
		"data": gin.H{
			"name": "About me",
			"age":  "99",
		},
	})
}
