package model

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	path := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       path,
		DefaultStringSize:         256,  // string字段默认长度
		DisableDatetimePrecision:  true, //禁止datatime精度，mysql5.6之前的数据库不支持
		DontSupportRenameIndex:    true, //禁止重命名索引， 即禁止把索引删了重建，5.7
		DontSupportRenameColumn:   true, //禁止重命名列，8之前不支持
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		//日志
		Logger: ormLogger,
		//命名策略，单数化不加s
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database ,err:" + err.Error())
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)                  //设置空闲连接池
	sqlDB.SetMaxOpenConns(100)                 //最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30) //限制连接的生命周期，连接超时后，会在需要时惰性回收复用
	DB = db
	//数据库迁移
	migration()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := DB
	return db.WithContext(ctx)
}

func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&User{},
			&Video{},
			&Review{},
			&Interactive{},
			&Comment{},
			&Reply{},
			&Announce{},
			&AnnounceUser{},
			&Messages{},
			&Danmu{},
			&Carousel{},
			&Admin{},
		)
	if err != nil {
		fmt.Println("err: ", err)
	}
	return
}
