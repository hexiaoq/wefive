package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"wefive/model"
	"wefive/response"
	"wefive/service"
	"wefive/util"
)

func AddProcessForBus(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	busId, err1 := strconv.Atoi(mMap["busId"])
	if err1 != nil {
		response.Fail(ctx, nil, err1.Error())
		return
	}
	var process model.Process
	process.Title = mMap["title"]
	process.Content = mMap["content"]
	process.Step, err1 = strconv.Atoi(mMap["step"])
	if err1 != nil {
		response.Fail(ctx, nil, err1.Error())
		return
	}
	process.BusId = int64(busId)
	if err := service.AddProcessForBus(&process); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "为业务添加流程成功！")
}

func AddProcessMaterial(ctx *gin.Context) {
	processId, err := strconv.Atoi(ctx.Param("processId"))
	if err != nil {
		response.Fail(ctx, nil, err.Error())
	}
	var infoMap = make(map[string]interface{})
	json.NewDecoder(ctx.Request.Body).Decode(&infoMap)
	materials := infoMap["processMaterials"]
	if err1 := service.AddProcessMaterials(int64(processId), materials); util.IsFailed(err1) {
		response.Fail(ctx, nil, err1.Message)
		return
	}

	response.Success(ctx, nil, "为流程添加材料成功！")
}

func SendAllProcessOfBus(ctx *gin.Context) {
	var err1 *util.Err
	var processes *[]model.ProcessDto
	busId, err := strconv.Atoi(ctx.Param("busId"))
	if err != nil {
		response.Fail(ctx, nil, err.Error())
		return
	}
	if processes, err1 = service.GetAllProcessOfBus(int64(busId)); util.IsFailed(err1) {
		response.Fail(ctx, nil, err1.Message)
		return
	}
	response.Success(ctx, gin.H{
		"processes": processes,
	}, "获取业务所有流程成功！")
}

/*func SendBusProcess(ctx *gin.Context) {
	busId, err := strconv.Atoi(ctx.Param("busId"))
	if err != nil {
		response.Fail(ctx, nil, err.Error())
		return
	}
}*/

func DeleteProcess(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	busId, err1 := strconv.Atoi(mMap["busId"])
	if err1 != nil {
		response.Fail(ctx, nil, err1.Error())
		return
	}
	if err := service.DeleteProcessByBusId(int64(busId)); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "删除业务流程成功！")
}

func DeleteProcessMaterial(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	processId, err1 := strconv.Atoi(mMap["processId"])
	if err1 != nil {
		response.Fail(ctx, nil, err1.Error())
		return
	}
	materialId, err1 := strconv.Atoi(mMap["materialId"])
	if err1 != nil {
		response.Fail(ctx, nil, err1.Error())
		return
	}
	if err := service.DeleteProcessMaterial(int64(processId), int64(materialId)); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "删除业务流程材料成功！")
}
