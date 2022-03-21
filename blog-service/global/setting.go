package global

import (
	"github.com/ydnju/go_tour/blog-service/pkg/logger"
	"github.com/ydnju/go_tour/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	Logger          *logger.Logger
)
