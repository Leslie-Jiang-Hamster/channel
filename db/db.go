package db

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var sqldb *sql.DB

func Connect(dbConfig gorm.Dialector, logger_ logger.Interface) {
	var err error;
	db, err := gorm.Open(dbConfig, &gorm.Config{
		Logger: logger_,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	sqldb, err = db.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetSQLDB() *sql.DB {
	return sqldb
}

