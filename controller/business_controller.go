package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gover-server/model"
	"gover-server/response"
	"gover-server/service"
	"gover-server/util"
	"log"
	"strconv"
)

func SendAllBusiness(ctx *gin.Context) {
	var businesses *[]model.Business
	var err *util.Err
	if businesses, err = service.GetAllBusinesses(); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, gin.H{
		"businesses": *businesses,
	}, "获取业务成功！")
	return
}

func SendAllBusinessOfDeptByDeptId(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	deptId, err1 := strconv.Atoi(mMap["deptId"])
	if err1 != nil {
		log.Println(err1)
		response.Fail(ctx, nil, err1.Error())
		return
	}
	var businesses *[]model.Business
	var err *util.Err
	if businesses, err = service.GetAllBusinessesByDeptId(int64(deptId)); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, gin.H{
		"businesses": *businesses,
	}, "获取业务成功！")
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

// 获取热门业务
func SendHotBusiness(ctx *gin.Context) {
	businesses, err := service.GetHotBusiness()
	if util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, gin.H{
		"businesses": *businesses,
	}, "获取热门业务成功！")
}

func AddBusiness(ctx *gin.Context) {
	var busMap = make(map[string]string)
	var business model.Business
	json.NewDecoder(ctx.Request.Body).Decode(&busMap)
	deptId, err1 := strconv.Atoi(ctx.Param("deptId"))
	if err1 != nil {
		response.Fail(ctx, nil, err1.Error())
		return
	}
	business.BusName = busMap["busName"]
	business.Description = busMap["description"]
	business.Requirement = busMap["requirement"]
	business.Cost, _ = strconv.ParseFloat(busMap["cost"], 8)
	if err := service.AddBusiness(&business, int64(deptId)); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "添加业务成功！")
}

func DeleteBusiness(ctx *gin.Context) {
	busId, err1 := strconv.Atoi(ctx.Param("busId"))
	if err1 != nil {
		response.Fail(ctx, nil, err1.Error())
		return
	}
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

// 从业务模板中添加业务
func AddBusTemplate(ctx *gin.Context) {
	var busMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&busMap)
	busName := busMap["busName"]
	deptId, err1 := strconv.Atoi(busMap["deptId"])
	if err1 != nil {
		response.Fail(ctx, nil, err1.Error())
		return
	}
	if err := service.AddBusTemplate(busName, int64(deptId)); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "添加业务模板成功！")
}
