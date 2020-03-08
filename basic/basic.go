package basic

import (
	"basic/config"
	"basic/db"
	"basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}