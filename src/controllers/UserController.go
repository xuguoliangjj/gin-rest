package controllers

import (
	"base"
	"database/sql"
	"db"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	base.Controller
}

type User struct {
	Id       int64          `db:"id"`
	UserName sql.NullString `db:"name"`
	Email    sql.NullString `db:"email"`
}

var GetUserController UserController

func (this *UserController) Index(ctx *gin.Context) {
	user := new(User)
	row := db.OBJ.DBHand.QueryRow("SELECT id,username,email FROM user where username = ?", "admin")
	err := row.Scan(&user.Id, &user.UserName, &user.Email)
	if err == nil {
		ctx.JSON(200, gin.H{
			"code":    1,
			"message": "ok",
			"data": gin.H{
				"name":  user.UserName.String,
				"email": user.Email.String,
			},
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "fail",
		})
	}
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
