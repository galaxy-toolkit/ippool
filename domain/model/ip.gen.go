// Package model query model
package model

const TableNameIP = "ip"

// IP mapped from table <ip>
type IP struct {
	ID        int64      `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:ID" json:"id"`                                   // ID
	Address   string     `gorm:"column:address;type:character varying;not null;default:''::character varying;comment:IP 地址" json:"address"`  // IP 地址
	Status    IPStatus   `gorm:"column:status;type:character varying;not null;default:''::character varying;comment:状态" json:"status"`       // 状态
	Latency   int64      `gorm:"column:latency;type:bigint;not null;comment:延迟" json:"latency"`                                              // 延迟
	Source    string     `gorm:"column:source;type:character varying;not null;default:''::character varying;comment:来源" json:"source"`       // 来源
	Port      string     `gorm:"column:port;type:character varying;not null;default:''::character varying;comment:端口" json:"port"`           // 端口
	Protocol  IPProtocol `gorm:"column:protocol;type:character varying;not null;default:''::character varying;comment:代理协议" json:"protocol"` // 代理协议
	Location  string     `gorm:"column:location;type:character varying;not null;default:''::character varying;comment:位置" json:"location"`   // 位置
	CrawlTime int64      `gorm:"column:crawl_time;type:bigint;not null;comment:爬取时间，时间戳" json:"crawl_time"`                                  // 爬取时间，时间戳
	CheckTime int64      `gorm:"column:check_time;type:bigint;not null;comment:验证时间，时间戳" json:"check_time"`                                  // 验证时间，时间戳
}

// TableName IP's table name
func (*IP) TableName() string {
	return TableNameIP
}
