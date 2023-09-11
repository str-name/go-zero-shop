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

const PaymentDSN = "root:gaj991130@tcp(127.0.0.1:3306)/zero_shop_payment?charset=utf8mb4&parseTime=True&loc=Local"

var PaymentDB *gorm.DB

func init() {
	newLogger := logger.New(log.New(os.Stdout, "zero-shop   ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Error,
		})
	db, err := gorm.Open(mysql.Open(PaymentDSN),
		&gorm.Config{
			Logger: newLogger,
		})
	if err != nil {
		logx.WithContext(context.Background()).Errorf("GORM connect PaymentDB Error: %+v", err)
	}
	PaymentDB = db
}
