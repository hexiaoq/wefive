package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"wefive/model"
	"wefive/response"
	"wefive/service"
	"wefive/util"
)

func SendAllGovernors(ctx *gin.Context) {
	var governors *[]model.GovernorDto
	var err *util.Err
	if governors, err = service.GetAllGovernors(); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, gin.H{
		"governors": *governors,
	}, "获取人员列表成功！")
}

func AddGovernor(ctx *gin.Context) {
	var goverMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&goverMap)
	phone := goverMap["phone"]
	deptName := goverMap["department"]
	if err := service.AddGovernor(phone, deptName); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "添加工作人员成功！")
}

func DeleteGovernor(ctx *gin.Context) {
	var goverMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&goverMap)
	phone := goverMap["phone"]
	if err := service.DeleteGovernorByPhone(phone); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "删除工作人员成功！")
}
