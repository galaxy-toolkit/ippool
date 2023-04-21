package global

import (
	"github.com/galaxy-toolkit/server/database/mysql"
	"gorm.io/gorm"
)

// MySQL 数据库链接
var MySQL *gorm.DB

// InitMySQL 初始化 MySQL 数据库
func InitMySQL() {
	db, err := mysql.New(Config.MySQL)
	if err != nil {
		panic(err)
	}
	MySQL = db
}
