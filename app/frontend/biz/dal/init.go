package dal

import (
	"github.com/CHlluanma/go-mall-kitex/app/frontend/biz/dal/mysql"
	"github.com/CHlluanma/go-mall-kitex/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
