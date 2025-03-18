package dal

import (
	"github.com/somewan820/tiktok/app/frontend/biz/dal/mysql"
	"github.com/somewan820/tiktok/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
