// Package model query model
package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID        int32          `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID" json:"id"`                  // ID
	Username  string         `gorm:"column:username;type:character varying;not null;comment:用户名，用于登录" json:"username"`           // 用户名，用于登录
	Password  string         `gorm:"column:password;type:character varying;not null;comment:密码" json:"password"`                 // 密码
	Birthday  time.Time      `gorm:"column:birthday;type:date;not null;comment:生日" json:"birthday"`                              // 生日
	Gender    string         `gorm:"column:gender;type:character varying;not null;comment:性别" json:"gender"`                     // 性别
	Avatar    string         `gorm:"column:avatar;type:character varying;not null;comment:头像" json:"avatar"`                     // 头像
	Email     string         `gorm:"column:email;type:character varying;not null;comment:邮箱" json:"email"`                       // 邮箱
	Phone     string         `gorm:"column:phone;type:character varying;not null;comment:手机号" json:"phone"`                      // 手机号
	Status    string         `gorm:"column:status;type:character varying;not null;comment:状态" json:"status"`                     // 状态
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp without time zone;not null;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp without time zone;not null;comment:更新事件" json:"updated_at"` // 更新事件
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp without time zone;comment:删除事件" json:"deleted_at"`          // 删除事件
	Nickname  string         `gorm:"column:nickname;type:character varying;not null;comment:昵称，用于展示" json:"nickname"`            // 昵称，用于展示
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
