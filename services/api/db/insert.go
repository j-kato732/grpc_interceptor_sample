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

	// if err = db.Create(&model.UserInfo{
	// 	UserId:        2,
	// 	LastName:      "伊藤",
	// 	FirstName:     "優",
	// 	Period:        "202105",
	// 	DepartmentId:  1,
	// 	JobId:         1,
	// 	EnrollmentFlg: true,
	// 	AdminFlg:      false,
	// }).Error; err != nil {
	// 	log.Println(err)
	// }
}
