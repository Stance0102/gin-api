package main

import (
	"fmt"
	"gin-api/global"
	"gin-api/internal/model"
	"gin-api/internal/pkg/setting"
)

func main() {
	setting, _ := setting.NewSetting()
	_ = setting.ReadSection("Database", &global.DatabaseSetting)

	db, err := model.NewDBEngine(global.DatabaseSetting)
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println(err)
	}
}
