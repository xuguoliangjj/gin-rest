package controllers

import (
	"base"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"models"
	"time"
)

type UserController struct {
	base.Controller
}

var GetUserController UserController

func (this *UserController) Index(ctx *gin.Context, p *base.SQLConnPool) {
	var users []*models.User
	username := "许国梁"
	rows, _ := p.SQLDB.Query("SELECT id,username,email FROM user where username = ?", username)
	defer rows.Close()
	for rows.Next() {
		user := new(models.User)
		err := rows.Scan(&user.Id, &user.UserName, &user.Email)
		if err != nil {
			panic(err)
		} else {
			users = append(users, user)
		}
	}
	var data []gin.H
	for _, row := range users {
		data = append(data, gin.H{
			"id":       row.Id,
			"username": row.UserName.String,
			"email":    row.Email.String,
		})
	}
	ctx.JSON(200, gin.H{
		"code":    1,
		"message": "ok",
		"data":    data,
	})
}

func doInsert(p *base.SQLConnPool, username string) int64 {
	lastId, err := p.Insert("insert into user (username,auth_key,password_hash,email) value (?,'1','1','1')", username)
	if err != nil {
		log.Panicf("错误：%s\n", err.Error())
		return 0
	}
	fmt.Println(time.Now(), lastId)
	return lastId
}
func (this *UserController) Create(ctx *gin.Context, p *base.SQLConnPool) {
	go doInsert(p, ctx.Param("username"))
	ctx.JSON(200, gin.H{
		"code":    1,
		"message": "ok",
		"data": gin.H{
			"name": "Create: " + ctx.Param("username"),
			"age":  "99",
		},
	})
}
