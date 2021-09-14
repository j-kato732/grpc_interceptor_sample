package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"grpc_gateway_sample/db/model"
	// pb "grpc_gateway_sample/proto"
)

const (
	conn = "host=db port=5432 user=admin password=password+1 dbname=testdb sslmode=disable TimeZone=Asia/Shanghai"
)

// var (
// 	periods   []model.Period
// 	userInfos model.UserInfo
// )

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	con, err := db.DB()
	defer con.Close()

	if err := db.Create(&model.Period{Period: "202105"}).Error; err != nil {
		log.Println(err)
	}
}
