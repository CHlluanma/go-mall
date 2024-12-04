package dal

import (
	"github.com/chhz0/go-mall-kitex/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
