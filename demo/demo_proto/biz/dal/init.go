package dal

import (
	"github.com/CHlluanma/go-mall-kitex/demo/demo_proto/biz/dal/mysql"
	"github.com/CHlluanma/go-mall-kitex/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
