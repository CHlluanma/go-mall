package dal

import (
	"github.com/CHlluanma/go-mall-kitex/app/cart/biz/dal/mysql"
	"github.com/CHlluanma/go-mall-kitex/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
