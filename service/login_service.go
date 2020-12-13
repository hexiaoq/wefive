package service

import (
	"log"
	"wefive/common"
	"wefive/model"
	"wefive/util"
)

// 验证登录账户和密码
func CheckLoginValidation(password string, phone string) (*model.GovernorDto, *util.Err) {
	db := common.GetDB()
	var gover model.Governor
	var goverDto *model.GovernorDto
	var err *util.Err
	db.Where("phone = ?", phone).First(&gover)
	if gover.GoverId == 0 {
		log.Println("用户不存在！phone = ", phone)
		return goverDto, util.Fail("用户不存在")
	}
	if password != gover.Password {
		log.Println("密码错误！phone = ", phone)
		return goverDto, util.Fail("密码错误")
	}
	goverDto, err = ToGovernorDto(&gover)
	if util.IsFailed(err) {
		log.Println(err.Message)
		return goverDto, util.Fail(err.Message)
	}
	return goverDto, util.Success()
}

// 验证登录账户和密码
func CheckAdminLoginValidation(password string, phone string) (model.Admin, *util.Err) {
	db := common.GetDB()
	var admin model.Admin
	db.Where("phone = ?", phone).First(&admin)
	if admin.AdminId == 0 {
		log.Println("管理员不存在！phone = ", phone)
		return admin, util.Fail("用户不存在")
	}
	if password != admin.Password {
		log.Println("管理员密码错误！phone = ", phone)
		return admin, util.Fail("密码错误")
	}
	return admin, util.Success()
}
