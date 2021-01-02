package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gover-server/model"
	"gover-server/response"
	"gover-server/service"
	"gover-server/util"
)

type Conf struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

// 验证登录的账号和密码
func LoginCheck(ctx *gin.Context) {
	var goverMap = make(map[string]string)
	var goverDto *model.GovernorDto
	var err *util.Err
	json.NewDecoder(ctx.Request.Body).Decode(&goverMap)
	password := goverMap["password"]
	phone := goverMap["phone"]
	if goverDto, err = service.CheckLoginValidation(password, phone); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, gin.H{
		"governor": goverDto,
	}, "欢迎登录！")
}

// 验证管理员登录的账号和密码
func LoginAdminCheck(ctx *gin.Context) {
	var adminMap = make(map[string]string)
	var admin model.Admin
	var err *util.Err
	json.NewDecoder(ctx.Request.Body).Decode(&adminMap)
	password := adminMap["password"]
	phone := adminMap["phone"]
	if admin, err = service.CheckAdminLoginValidation(password, phone); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, gin.H{
		"admin": admin,
	}, "欢迎登录！")
}
