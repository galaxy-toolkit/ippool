package global

import (
	"github.com/galaxy-toolkit/server/database/mysql"
	"gorm.io/gen"
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

func InitMySQLGenerator(queryPath, queryFilename, modelPath string) (*gen.Generator, error) {
	generator, err := mysql.NewModelGenerator(Config.MySQL, queryPath, queryFilename, modelPath)
	if err != nil {
		return nil, err
	}
	return generator, nil
}
