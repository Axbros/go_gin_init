package global

import (
	"gin_init/pkg/logger"
	"gin_init/pkg/setting"
)

var (
	Logger          *logger.Logger
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
)
