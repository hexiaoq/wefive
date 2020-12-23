package service

import (
	"log"
	"wefive/common"
	"wefive/model"
	"wefive/util"
)

func GetUserByUserId(userId int64) (*model.User, *util.Err) {
	var user model.User
	db := common.GetDB()
	err := db.Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}
	return &user, util.Success()
}
