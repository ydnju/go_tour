package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ydnju/go_tour/blog-service/global"
	"github.com/ydnju/go_tour/blog-service/internal/model"
	"github.com/ydnju/go_tour/blog-service/internal/routers"
	"github.com/ydnju/go_tour/blog-service/pkg/logger"
	"github.com/ydnju/go_tour/blog-service/pkg/setting"
)

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error

	// 用一个global变量来存储global的内容，比如配置对象、数据库引擎等
	// 注意这里不能用 :=，因为 :=会重新声明并创建左侧的新局部变量，调用完成后，
	// global.DBEngine对象仍然是nil，没有真正赋值到包的全局变量global.DBEngine上面去
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupLogger() error {

	/*
		global.Logger = logger.NewLogger(&lumberjack.Logger{
			Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
			MaxSize:   600,
			MaxAge:    10,
			LocalTime: true,
		}, "", log.LstdFlags).WithCaller(2)
	*/

	// In our case, just write logs to stdout to be
	// collected by kubernetes
	global.Logger = logger.NewLogger(os.Stdout, "", log.LstdFlags).WithCaller(2)

	return nil
}

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/ydnju/go_tour
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	global.Logger.Infof("%s: go_tour/%s", "eddycjy", "blog-service")

	s.ListenAndServe()
}
