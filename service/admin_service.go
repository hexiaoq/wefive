package service

import (
	"github.com/spf13/viper"
	"log"
	"wefive/common"
	"wefive/model"
)

// 在数据库中判断一名申请登录的用户是否为管理员
func IsAdminExist(phone string) bool {
	db := common.GetDB()
	var admin model.Admin
	log.Println("待查找的管理员账号：", phone)
	db.Where("phone = ?", phone).First(&admin)
	log.Println("找到的管理员ID：", admin.AdminId)
	if admin.AdminId != 0 {
		return true
	}
	return false
}

// 在数据库中初始化一名高级管理员
func InitAdmin() {
	DB := common.GetDB()
	admin := model.Admin{
		Password: viper.GetString("admin.password"),
		Name:     viper.GetString("admin.name"),
		Phone:    viper.GetString("admin.phone"),
	}

	if IsAdminExist(admin.Phone) {
		log.Println("高级用户已经存在！")
	} else {
		log.Println("高级用户不存在！已经创建")
		DB.Create(&admin)
	}
}
