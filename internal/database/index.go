package database

import (
	"log"

	"github.com/verlinof/softlancer-go/internal/config/db_config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	dsnMysql := db_config.DB_USER + ":" + db_config.DB_PASSWORD + "@tcp(" + db_config.DB_HOST + ":" + db_config.DB_PORT + ")/" + db_config.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsnMysql), &gorm.Config{})

	//========Config Postgres========
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// dsn := "host=" + db_config.DB_HOST + " user=" + db_config.DB_USER + " password=" + db_config.DB_PASSWORD + " dbname=" + db_config.DB_NAME + " port=" + db_config.DB_PORT + " sslmode=disable TimeZone=Asia/Shanghai"
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	log.Println("Database connected!")
}
