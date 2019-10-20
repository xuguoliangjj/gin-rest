package base

type DBModel struct {
}

var MODEL DBModel

func (this *DBModel) init() {
	this.InitDB()
}

func (this *DBModel) InitDB() bool {
	OBJ.DBHand = OBJ.GetMySQL()
	//db.OBJ.DBHand.SetMaxOpenConns(100)
	//db.OBJ.DBHand.SetMaxIdleConns(100)
	//db.OBJ.DBHand.SetConnMaxLifetime(10)
	err := OBJ.DBHand.Ping()
	return err == nil
}
