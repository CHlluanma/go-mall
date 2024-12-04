package dal

import (
	"github.com/chhz0/go-mall-kitex/demo/demo_thrift/biz/dal/mysql"
	"github.com/chhz0/go-mall-kitex/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
