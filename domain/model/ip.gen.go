// Package model query model
package model

const TableNameIP = "ip"

// IP mapped from table <ip>
type IP struct {
	ID      int64    `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:ID" json:"id"`    // ID
	Address string   `gorm:"column:address;type:character varying;not null;comment:IP 地址" json:"address"` // IP 地址
	Status  IPStatus `gorm:"column:status;type:character varying;not null;comment:状态" json:"status"`      // 状态
	Latency int64    `gorm:"column:latency;type:bigint;not null;comment:延迟" json:"latency"`               // 延迟
	Source  string   `gorm:"column:source;type:character varying;not null;comment:来源" json:"source"`      // 来源
}

// TableName IP's table name
func (*IP) TableName() string {
	return TableNameIP
}
