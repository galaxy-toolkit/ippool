// Package model query model
package model

import (
	"time"
)

const TableNameIP = "ip"

// IP mapped from table <ip>
type IP struct {
	ID        int64      `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:ID" json:"id"`                                           // ID
	Address   string     `gorm:"column:address;type:character varying;not null;uniqueIndex:un_address_port,priority:1;comment:IP 地址" json:"address"` // IP 地址
	Status    IPStatus   `gorm:"column:status;type:character varying;not null;comment:状态" json:"status"`                                             // 状态
	Latency   int64      `gorm:"column:latency;type:bigint;not null;comment:延迟" json:"latency"`                                                      // 延迟
	Source    string     `gorm:"column:source;type:character varying;not null;comment:来源" json:"source"`                                             // 来源
	Port      string     `gorm:"column:port;type:character varying;not null;uniqueIndex:un_address_port,priority:2;comment:端口" json:"port"`          // 端口
	Protocol  IPProtocol `gorm:"column:protocol;type:character varying;not null;comment:代理协议" json:"protocol"`                                       // 代理协议
	Location  string     `gorm:"column:location;type:character varying;not null;comment:位置" json:"location"`                                         // 位置
	CreateAt  time.Time  `gorm:"column:create_at;type:timestamp without time zone;not null;comment:创建时间" json:"create_at"`                           // 创建时间
	CheckTime time.Time  `gorm:"column:check_time;type:timestamp without time zone;not null;comment:验证时间" json:"check_time"`                         // 验证时间
}

// TableName IP's table name
func (IP) TableName() string {
	return TableNameIP
}
