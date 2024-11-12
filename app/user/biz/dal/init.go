package dal

import (
	"github.com/CHlluanma/go-mall-kitex/app/user/biz/dal/mysql"
	"github.com/CHlluanma/go-mall-kitex/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
