package dal

import (
	"github.com/CHlluanma/go-mall-kitex/app/checkout/biz/dal/mysql"
	"github.com/CHlluanma/go-mall-kitex/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
