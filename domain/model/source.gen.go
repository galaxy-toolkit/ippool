// Package model query model
package model

const TableNameSource = "source"

// Source mapped from table <source>
type Source struct {
	ID      int32  `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID" json:"id"`  // ID
	Name    string `gorm:"column:name;type:character varying;not null;comment:来源名称" json:"name"`       // 来源名称
	Website string `gorm:"column:website;type:character varying;not null;comment:来源站点" json:"website"` // 来源站点
}

// TableName Source's table name
func (*Source) TableName() string {
	return TableNameSource
}
