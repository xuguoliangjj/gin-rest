package base

import (
	"log"
	"os"
	"path/filepath"
)

type InitSt struct {
	cfg Config
}

var INIT_OBJ InitSt

//加载配置文件
func (this *InitSt) Init() bool {
	dir, e := filepath.Abs(filepath.Dir(os.Args[0]))
	if e != nil {
		log.Fatal(e)
	}
	dir += "/server.cfg"
	this.cfg.Read(dir)
	return true
}
