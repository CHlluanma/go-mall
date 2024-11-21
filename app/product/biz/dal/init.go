package dal

import (
	"github.com/CHlluanma/go-mall-kitex/app/product/biz/dal/mysql"
	"github.com/CHlluanma/go-mall-kitex/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
