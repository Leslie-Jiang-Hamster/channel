package bootstrap

import (
	"fmt"
	"time"

	"github.com/channel/db"
	"github.com/channel/env"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDB(env env.IEnv) {
	var dbConfig gorm.Dialector
	switch env.DBEnv.Connection {
	case "mysql":
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
		env.DBEnv.Username,
		env.DBEnv.Password,
		env.DBEnv.Host,
		env.DBEnv.Port,
		env.DBEnv.Database,
		env.DBEnv.Charset,
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqlite":
		dbConfig = sqlite.Open(env.DBEnv.Database)
	default:
		panic(fmt.Errorf("db connection not supported"))
	}
	db.Connect(dbConfig, logger.Default.LogMode(logger.Info))
	db.GetSQLDB().SetMaxOpenConns(env.DBEnv.Max_open_conns)
	db.GetSQLDB().SetMaxIdleConns(env.DBEnv.Max_idle_conns)
	db.GetSQLDB().SetConnMaxLifetime(time.Duration(env.DBEnv.Max_life_seconds) * time.Second)
}