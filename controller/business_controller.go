package controller

import (
	"claps-admin/model"
	"claps-admin/response"
	"claps-admin/service"
	"claps-admin/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

func SendAllBusinessOfDept(ctx *gin.Context) {
	deptName := ctx.Param("name")
	var businesses *[]model.Business
	var err *util.Err
	if businesses, err = service.GetAllBusinessesByDeptName(deptName); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, gin.H{
		"businesses": *businesses,
	}, "获取业务成功！")
	return
}

func AddBusiness(ctx *gin.Context) {
	var busMap = make(map[string]string)
	var business model.Business
	json.NewDecoder(ctx.Request.Body).Decode(&busMap)
	deptName := ctx.Param("name")
	business.BusName = busMap["busName"]
	business.Description = busMap["description"]
	business.Requirement = busMap["requirement"]
	business.Cost, _ = strconv.ParseFloat(busMap["cost"], 8)
	if err := service.AddBusiness(&business, deptName); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "添加业务成功！")
}

func DeleteBusiness(ctx *gin.Context) {
	//deptName := ctx.Param("name")
	var busMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&busMap)
	busId, _ := strconv.Atoi(busMap["busId"])
	if err := service.DeleteBusinessById(int64(busId)); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "删除业务成功！")
}
