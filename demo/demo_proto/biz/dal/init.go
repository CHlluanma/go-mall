package dal

import (
	"github.com/CHlluanma/go-mall-kitex/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
