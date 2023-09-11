package gorm

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const OrderDSN = "root:123456@tcp(127.0.0.1:3306)/zero_shop_order?charset=utf8mb4&parseTime=True&loc=Local"

var OrderDB *gorm.DB

func init() {
	newLogger := logger.New(log.New(os.Stdout, "zero-shop   ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Error,
		})
	db, err := gorm.Open(mysql.Open(OrderDSN),
		&gorm.Config{
			Logger: newLogger,
		})
	if err != nil {
		logx.WithContext(context.Background()).Errorf("GORM connect OrderDB Error: %+v", err)
	}
	OrderDB = db
}
