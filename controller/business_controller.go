package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"wefive/model"
	"wefive/response"
	"wefive/service"
	"wefive/util"
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

func SendBusiness(ctx *gin.Context) {
	var busMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&busMap)
	busId, _ := strconv.Atoi(busMap["busId"])
	var business *model.Business
	var err *util.Err
	if business, err = service.GetBusinessById(int64(busId)); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, gin.H{
		"business": *business,
	}, "获取业务成功！")
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
	if err := service.DeleteBusiness(int64(busId)); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "删除业务成功！")
}

func UpdateBusiness(ctx *gin.Context) {
	var busMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&busMap)
	busId, cerr := strconv.Atoi(busMap["busId"])
	if cerr != nil {
		log.Println(cerr.Error())
		response.Fail(ctx, nil, cerr.Error())
		return
	}
	deptId, cerr := strconv.Atoi(busMap["deptId"])
	if cerr != nil {
		log.Println(cerr.Error())
		response.Fail(ctx, nil, cerr.Error())
		return
	}
	cost, cerr := strconv.ParseFloat(busMap["cost"], 64)
	if cerr != nil {
		log.Println(cerr.Error())
		response.Fail(ctx, nil, cerr.Error())
		return
	}
	var business model.Business
	business.BusId = int64(busId)
	business.DeptId = int64(deptId)
	business.Description = busMap["description"]
	business.Requirement = busMap["requirement"]
	business.BusName = busMap["busName"]
	business.Cost = cost

	if err := service.UpdateBusiness(&business); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "业务修改成功！")
}
