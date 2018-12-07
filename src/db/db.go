package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DBObj struct {
	DBHand *sql.DB
}

var OBJ DBObj

func (this *DBObj) OpenDB(svr string, usr string, pwd string, db string) *sql.DB {
	str := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", usr, pwd, svr, db)
	hand, err := sql.Open("mysql", str)
	if err != nil {
		log.Panicf("错误：%s\n", err.Error())
	}
	return hand
}

func (this *DBObj) InitDB() bool {
	DB_Server := "127.0.0.1"
	DB_UserId := "root"
	DB_Password := "123456"
	DB_Name := "admin"
	this.DBHand = this.OpenDB(DB_Server, DB_UserId, DB_Password, DB_Name)
	err := this.DBHand.Ping()
	return err != nil
}
