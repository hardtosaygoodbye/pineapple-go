package database

import (
	"pineapple-go/config"
	"pineapple-go/core/log"
	"pineapple-go/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

// ConnectDB connect to database
func ConnectDB() {
	databaseURL := config.DB.DATABASE_URL
	var err error
	logger := log.DBLogger{}

	mysqlConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.DB.PREFIX, // table name prefix, table for `User` would be `t_users`
			SingularTable: true,             // use singular table name, table for `User` would be `user` with this option enabled
		},
		Logger: &logger,
	}
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       databaseURL, // data source name
		DefaultStringSize:         255,         // default size for string fields
		DisableDatetimePrecision:  true,        // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,        // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,        // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,       // auto configure based on currently MySQL version
	}), &mysqlConfig)
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	//设置与数据库建立连接的最大数目
	sqlDB.SetMaxOpenConns(config.DB.MaxOpenConns)
	//设置连接池中的最大闲置连接数
	sqlDB.SetMaxIdleConns(config.DB.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.DB.ConnMaxLifeTime)

	log.InitLogger.Info("connnect to mysql database successful")

}

//DisconnectDB disconnect database
func DisconnectDB() {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		panic(err)
	}
}

// DB inject with ctx for log
func DB(ctx *gin.Context) *gorm.DB {
	return db.WithContext(ctx)
}

//AutoMigrate auto migrate table to database
func AutoMigrate() {

	//设置表默认属性
	tableoptions := "CHARSET=" + config.DB.CHARSET

	log.GormLogger.Info("migrate start...")
	db.Set("gorm:table_options", tableoptions).AutoMigrate(
		&model.User{},
		&model.AuthCode{},
		&model.Weico{},
		&model.WeicoPic{},
		&model.WeicoLike{},
		&model.WeicoComment{},
		&model.WeicoCate{},
	)
	log.GormLogger.Info("migrate end...")
	log.InitLogger.Info("migrate table successful")
}
