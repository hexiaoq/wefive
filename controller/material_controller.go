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

func SendMaterials(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	busId, err1 := strconv.Atoi(mMap["busId"])
	if err1 != nil {
		response.Fail(ctx, nil, err1.Error())
		return
	}
	var materials *[]model.Material
	var err *util.Err
	if materials, err = service.GetMaterialsByBusId(int64(busId)); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, gin.H{
		"materials": *materials,
	}, "获取业务材料成功！")
}

func AddMaterial(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	busId, _ := strconv.Atoi(mMap["busId"])
	var material model.Material
	material.MaterialName = mMap["materialName"]
	material.Description = mMap["description"]
	material.PhotoUrl = mMap["photoUrl"]
	if err := service.AddMaterial(&material, int64(busId)); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "材料添加成功！")
}

func UpdateMaterial(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	busId, cerr := strconv.Atoi(mMap["busId"])
	if cerr != nil {
		log.Println(cerr.Error())
		response.Fail(ctx, nil, cerr.Error())
		return
	}
	materialId, cerr := strconv.Atoi(mMap["materialId"])
	if cerr != nil {
		log.Println(cerr.Error())
		response.Fail(ctx, nil, cerr.Error())
		return
	}
	var material model.Material
	material.MaterialId = int64(materialId)
	material.BusId = int64(busId)
	material.MaterialName = mMap["materialName"]
	material.Description = mMap["description"]
	material.PhotoUrl = mMap["photoUrl"]

	if err := service.UpdateMaterial(&material); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "材料修改成功！")
}

func DeleteMaterial(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	materialId, cerr := strconv.Atoi(mMap["materialId"])
	if cerr != nil {
		response.Fail(ctx, nil, cerr.Error())
		return
	}
	if err := service.DeleteMaterial(int64(materialId)); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "删除业务材料成功！")
}
