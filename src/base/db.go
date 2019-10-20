package base

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DBObj struct {
	DBHand *sql.DB
}

var OBJ DBObj

var MAX_POOL_SIZE = 100
var MySQLPool chan *sql.DB

func (this *DBObj) GetMySQL() *sql.DB {
	svr := INIT_OBJ.Cfg.Get("DB_Server")
	db := INIT_OBJ.Cfg.Get("DB_Name")
	usr := INIT_OBJ.Cfg.Get("DB_UserId")
	pwd := INIT_OBJ.Cfg.Get("DB_Password")
	if MySQLPool == nil {
		MySQLPool = make(chan *sql.DB, MAX_POOL_SIZE)
	}
	if len(MySQLPool) == 0 {
		go func() {
			for i := 0; i < MAX_POOL_SIZE/2; i++ {
				str := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", usr, pwd, svr, db)
				conn, err := sql.Open("mysql", str)
				if err != nil {
					panic(err)
				}
				putMySQL(conn)
			}
		}()
	}
	return <-MySQLPool
}

func putMySQL(conn *sql.DB) {
	if MySQLPool == nil {
		MySQLPool = make(chan *sql.DB, MAX_POOL_SIZE)
	}

	if len(MySQLPool) == MAX_POOL_SIZE {
		err := conn.Close()
		if err != nil {
			return
		}
		return
	}
	MySQLPool <- conn
}
