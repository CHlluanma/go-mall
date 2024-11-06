package main

import (
	"fmt"

	"github.com/CHlluanma/go-mall-kitex/demo/demo_proto/biz/dal"
	"github.com/CHlluanma/go-mall-kitex/demo/demo_proto/biz/dal/model"
	"github.com/CHlluanma/go-mall-kitex/demo/demo_proto/biz/dal/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()

	// Create
	// mysql.DB.Create(&model.User{Email: "demo@example.com", Password: "jfiojffjsoij"})

	// Read
	var row model.User
	mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").First(&row)

	fmt.Printf("row: %+v\n", row)
}
