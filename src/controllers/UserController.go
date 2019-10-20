package controllers

import (
	"base"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
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
	username := "许国梁"
	row := base.OBJ.DBHand.QueryRow("SELECT id,username,email FROM user where username = ?", username)
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
			"message": "not found: " + username,
		})
	}
}

func doInsert(ctx *gin.Context, username string) int64 {
	dbx := base.OBJ.GetMySQL()
	row, err := dbx.Exec("insert into user (username,auth_key,password_hash,email) value (?,'1','1','1')", username)
	defer dbx.Close()
	if row == nil {
		log.Panicf("错误：%s\n", err.Error())
		return 0
	}
	num, _ := row.LastInsertId()
	fmt.Println(time.Now(), num)
	return num
}
func (this *UserController) Create(ctx *gin.Context) {
	go doInsert(ctx, ctx.Param("username"))
	ctx.JSON(200, gin.H{
		"code":    1,
		"message": "ok",
		"data": gin.H{
			"name": "Create: " + ctx.Param("username"),
			"age":  "99",
		},
	})
}
