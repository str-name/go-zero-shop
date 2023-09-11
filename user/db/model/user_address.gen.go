// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserAddress = "user_address"

// UserAddress mapped from table <user_address>
type UserAddress struct {
	ID            int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	CreateTime    time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
	DeleteTime    time.Time `gorm:"column:delete_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"delete_time"`
	DelState      int64     `gorm:"column:del_state;type:tinyint;not null" json:"del_state"`
	UserID        int64     `gorm:"column:user_id;type:bigint;not null;index:ix_user_id,priority:1" json:"user_id"`
	IsDefault     int64     `gorm:"column:is_default;type:tinyint(1);not null;comment:是否为默认地址" json:"is_default"`
	Province      string    `gorm:"column:province;type:varchar(100);not null;comment:收货省份" json:"province"`
	City          string    `gorm:"column:city;type:varchar(100);not null;comment:收货城市" json:"city"`
	Region        string    `gorm:"column:region;type:varchar(100);not null;comment:收货区/县" json:"region"`
	DetailAddress string    `gorm:"column:detail_address;type:varchar(255);not null;comment:收货详情地址" json:"detail_address"`
	Name          string    `gorm:"column:name;type:varchar(64);not null;comment:收货人名称" json:"name"`
	Phone         string    `gorm:"column:phone;type:varchar(11);not null;comment:收货人手机号" json:"phone"`
}

// TableName UserAddress's table name
func (*UserAddress) TableName() string {
	return TableNameUserAddress
}