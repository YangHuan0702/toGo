package repository

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConn() *gorm.DB {
	dns := "host=120.76.140.220 user=postgres password=yanghuan dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error:{}", err)
	}
	return db
}
