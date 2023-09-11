// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameProductRecommend = "product_recommend"

// ProductRecommend mapped from table <product_recommend>
type ProductRecommend struct {
	ID         int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;not null;index:ix_update_time,priority:1;default:CURRENT_TIMESTAMP" json:"update_time"`
	DeleteTime time.Time `gorm:"column:delete_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"delete_time"`
	DelState   int64     `gorm:"column:del_state;type:tinyint;not null" json:"del_state"`
	ProductID  int64     `gorm:"column:product_id;type:bigint unsigned;not null;comment:商品id" json:"product_id"`
}

// TableName ProductRecommend's table name
func (*ProductRecommend) TableName() string {
	return TableNameProductRecommend
}
