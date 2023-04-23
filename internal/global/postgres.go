package global

import (
	"github.com/galaxy-toolkit/server/database/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Postgres 数据库链接
var Postgres *gorm.DB

// InitPostgres 初始化 Postgres 数据库
func InitPostgres() {
	db, err := postgres.New(Config.Database.Postgres)
	if err != nil {
		panic(err)
	}
	Postgres = db
}

func InitPostgresGenerator(gConf postgres.GeneratorConfig) (*gen.Generator, error) {
	generator, err := postgres.NewGenerator(Config.Database.Postgres, gConf)
	if err != nil {
		return nil, err
	}
	return generator, nil
}
