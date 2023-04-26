package pool

import (
	"context"
	"database/sql"

	"github.com/galaxy-toolkit/ippool/internal/global"
	"gorm.io/gorm"
)

type Query struct {
	DB *gorm.DB
	IP IPDao
}

func Use(ctx context.Context) *Query {
	db := global.Postgres.Session(&gorm.Session{
		Context: ctx,
	})

	return &Query{
		DB: db,
		IP: NewIPDao(ctx, db),
	}
}

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		DB: db,
		IP: q.IP.clone(db),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.DB.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}
