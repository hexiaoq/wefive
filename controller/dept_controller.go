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

func AddDepartment(ctx *gin.Context) {
	var deptMap = make(map[string]string)
	var department model.Department
	var err *util.Err
	json.NewDecoder(ctx.Request.Body).Decode(&deptMap)
	department.DeptName = deptMap["deptName"]
	department.Location = deptMap["location"]
	department.WorkTime = deptMap["wordTime"]
	department.Description = deptMap["description"]
	if err = service.AddDepartment(&department); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "添加部门成功！")
}

func DepartmentDelete(ctx *gin.Context) {
	var deptMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&deptMap)
	deptId, _ := strconv.Atoi(deptMap["deptId"])
	if err := service.DeleteDepartmentById(int64(deptId)); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "删除部门成功！")
}

func SendAllDepartments(ctx *gin.Context) {
	var departments *[]model.Department
	var err *util.Err
	if departments, err = service.GetAllDepartments(); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, gin.H{
		"departments": *departments,
	}, "获取部门列表成功！")
}
