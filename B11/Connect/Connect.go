package Connect

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=1234 dbname=Admins port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connect fail")
	} else {
		fmt.Print("Connect successfully")
	}
	DB = db
}
