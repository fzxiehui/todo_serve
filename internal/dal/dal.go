package dal

import (
	"fmt"
	"sync"
	"time"

	"github.com/fzxiehui/todo_serve/config"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB
var once sync.Once

func init() {
	once.Do(func() {
		err := Connect()
		if err != nil {
			panic(err)
		}
	})
}

func Connect() error {
	var err error
	dbuser := config.Config().GetString("db.user")
	dbpassword := config.Config().GetString("db.password")
	dbaddr := config.Config().GetString("db.addr")
	dbname := config.Config().GetString("db.name")

	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbuser, dbpassword, dbaddr, dbname)
	fmt.Println(dsn)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db, err := DB.DB()
	if err != nil {
		return err
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	db.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.SetConnMaxLifetime(time.Hour)

	return nil
}

func GetDB() *gorm.DB {
	return DB
}

func Close() {
	db, err := DB.DB()
	if err != nil {
		return
	}
	db.Close()
}
