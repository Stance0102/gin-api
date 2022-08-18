package main

import (
	"fmt"
	"gin-api/global"
	"gin-api/internal/model"
	"gin-api/internal/pkg/setting"
	"net/http"
	"time"

	routers "gin-api/internal"

	"github.com/gin-gonic/gin"
)

func init() {
	err := setupSetting()
	if err != nil {
		fmt.Printf("init.setupSetting error: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		fmt.Printf("init.setupDBEngine error: %v", err)
	}
}

func setupDBEngine() error {
	var err error
	global.DBEnging, err = model.NewDBEngine(global.DatabaseSetting)

	return err
}

func setupSetting() error {
	setting, err := setting.NewSetting()

	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)

	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)

	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	s := &http.Server{
		Addr:         ":" + global.ServerSetting.HttpPort,
		Handler:      router,
		ReadTimeout:  global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
	}

	s.ListenAndServe()
}
