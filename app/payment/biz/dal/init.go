package dal

import (
	"github.com/chhz0/go-mall-kitex/app/payment/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
