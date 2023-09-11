// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserHeaderImage = "user_header_image"

// UserHeaderImage mapped from table <user_header_image>
type UserHeaderImage struct {
	ID         int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
	DeleleTime time.Time `gorm:"column:delele_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"delele_time"`
	DelState   int64     `gorm:"column:del_state;type:tinyint;not null" json:"del_state"`
	Path       string    `gorm:"column:path;type:varchar(255);not null;comment:图片路径" json:"path"`
	Type       int64     `gorm:"column:type;type:tinyint(1);not null;default:2;comment:图片存储类型：1-本地存储，2-七牛云存储" json:"type"`
	Hash       *string   `gorm:"column:hash;type:longtext;comment:图片hash值，防止重复图片上传" json:"hash"`
}

// TableName UserHeaderImage's table name
func (*UserHeaderImage) TableName() string {
	return TableNameUserHeaderImage
}
