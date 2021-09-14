package main

import (
	"fmt"
	"log"
	// "reflect"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"grpc_gateway_sample/db/model"
	// pb "grpc_gateway_sample/proto"
)

var (
	periods   []model.Period
	userInfos model.UserInfo
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("faild to connect databse")
	}

	db.AutoMigrate(model.Period{})

	db.AutoMigrate(model.UserInfo{})

	fmt.Println("Success connected")
}
