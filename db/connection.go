package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var db *gorm.DB
var dbPool *sync.Pool

func Initialize() {

	dbStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", viper.GetString("MYSQL_USER"), viper.GetString("MYSQL_PASSWORD"), viper.GetString("MYSQL_HOST"), viper.GetString("MYSQL_PORT"), viper.GetString("MYSQL_DB"))

	//dsn := "user:password@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dbStr), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.Table("merchant_staff").Model(&MerchantStaff{})
	dbPool = &sync.Pool{
		New: func() interface{} {
			return db.Begin()
		},
	}
}

func GetDBConnection() *gorm.DB {
	conn := dbPool.Get().(*gorm.DB)
	return conn
}

func ReleaseDBConnection(conn *gorm.DB) {
	dbPool.Put(conn)
}
