package dal

import (
	"github.com/CHlluanma/go-mall-kitex/app/payment/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
