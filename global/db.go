package global

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

var (
	DBEngine  *gorm.DB
	RedisConn *redis.Pool
)
