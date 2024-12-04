package dal

import (
	"github.com/chhz0/go-mall-kitex/app/frontend/biz/dal/mysql"
	"github.com/chhz0/go-mall-kitex/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
