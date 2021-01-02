package service

import (
	"gover-server/common"
	"gover-server/model"
	"gover-server/util"
	"log"
)

func GetUserByUserId(userId int64) (*model.Users, *util.Err) {
	var user model.Users
	db := common.GetDB()
	err := db.Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}
	return &user, util.Success()
}
