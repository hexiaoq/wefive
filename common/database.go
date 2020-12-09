package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	//配置项读取
	driverName := viper.GetString("datasource.driverName")

	args := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True&loc=Local",
		viper.GetString("datasource.username"),
		viper.GetString("datasource.password"),
		viper.GetString("datasource.database"),
		viper.GetString("datasource.charset"))

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.SingularTable(true)
	DB = db
	/*// 数据库初始化表格
	if !db.HasTable(&model.User{}) {
		db.AutoMigrate(&model.User{})
	}
	if !db.HasTable(&model.Admin{}) {
		db.AutoMigrate(&model.Admin{})
	}
	if !db.HasTable(&model.Project{}) {
		db.AutoMigrate(&model.Project{})
	}*/
	return db
}

func GetDB() *gorm.DB {
	return DB
}
