package dal

import (
	"github.com/chhz0/go-mall-kitex/app/product/biz/dal/mysql"
	"github.com/chhz0/go-mall-kitex/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
