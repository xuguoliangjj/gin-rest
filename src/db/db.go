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
