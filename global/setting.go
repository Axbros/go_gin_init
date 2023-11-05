package global

import (
	"gin_init/pkg/logger"
	"gin_init/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	Logger          *logger.Logger
	RedisSetting    *setting.RedisSettingS
)
