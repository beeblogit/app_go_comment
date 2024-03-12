package bootstrap

import (
	"fmt"
	"log"
	"os"

	blogDomain "github.com/beeblogit/lib_go_domain/domain/blog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnection() (*gorm.DB, error) {

	dsn := os.ExpandEnv("$DATABASE_USER:$DATABASE_PASSWORD@($DATABASE_HOST:$DATABASE_PORT)/$DATABASE_NAME?charset=utf8&parseTime=True&loc=Local")
	fmt.Println(dsn)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if os.Getenv("DATABASE_DEBUG") == "true" {
		db = db.Debug()
	}

	if os.Getenv("DATABASE_MIGRATE") == "true" {
		if err := db.AutoMigrate(&blogDomain.Comment{}); err != nil {
			return nil, err
		}
	}

	return db, nil
}

func InitLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}
