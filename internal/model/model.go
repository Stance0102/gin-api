package model

import (
	"fmt"
	"gin-api/internal/pkg/setting"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBEngine(databaseSetting *setting.DatabaseSetting) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=gorm password=gorm dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		// databaseSetting.UserName,
		// databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Port,
		// databaseSetting.Charset,
		// databaseSetting.ParseTime,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
