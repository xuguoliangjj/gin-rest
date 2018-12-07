package base

import (
	"db"
)

type DBModel struct {
}

var MODEL DBModel

func (this *DBModel) init() {
	this.InitDB()
}

func (this *DBModel) InitDB() bool {
	DB_Server := INIT_OBJ.cfg.Get("DB_Server")
	DB_Name := INIT_OBJ.cfg.Get("DB_Name")
	DB_UserId := INIT_OBJ.cfg.Get("DB_UserId")
	DB_Password := INIT_OBJ.cfg.Get("DB_Password")

	db.OBJ.DBHand = db.OBJ.OpenDB(DB_Server, DB_UserId, DB_Password, DB_Name)
	db.OBJ.DBHand.SetMaxOpenConns(100)
	db.OBJ.DBHand.SetMaxIdleConns(100)
	db.OBJ.DBHand.SetConnMaxLifetime(10)
	err := db.OBJ.DBHand.Ping()
	return err == nil
}
