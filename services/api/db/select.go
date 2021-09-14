package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"grpc_gateway_sample/db/model"
)

const (
	conn = "host=db port=5432 user=admin password=password+1 dbname=testdb sslmode=disable TimeZone=Asia/Shanghai"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	con, err := db.DB()
	defer con.Close()

	var periods []model.Period

	// select * from Periods;
	if err := db.Find(&periods).Error; err != nil {
		fmt.Println(err)
	}

	fmt.Println(periods)

	var userInfo model.UserInfo
	// if err = db.Where("user_id = ? AND period = ?", "2", "202105").Find(&userInfo).Error; err != nil {
	// 	fmt.Println(err)
	// }
	if err = db.Where(model.UserInfo{UserId: 1, Period: "202105"}).Find(&userInfo).Error; err != nil {
		fmt.Println(err)
	}

	fmt.Println(userInfo)
}
