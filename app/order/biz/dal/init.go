package dal

import (
	"github.com/CHlluanma/go-mall-kitex/app/order/biz/dal/mysql"
	"github.com/CHlluanma/go-mall-kitex/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
