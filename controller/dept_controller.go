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

func AddDepartment(ctx *gin.Context) {
	var deptMap = make(map[string]string)
	var department model.Department
	var err *util.Err
	json.NewDecoder(ctx.Request.Body).Decode(&deptMap)
	department.DeptName = deptMap["deptName"]
	department.Location = deptMap["location"]
	department.WorkTime = deptMap["workTime"]
	department.Description = deptMap["description"]
	department.Picture = deptMap["picture"]
	department.Contact = deptMap["contact"]
	if err = service.AddDepartment(&department); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "添加部门成功！")
}

func DepartmentDelete(ctx *gin.Context) {
	var deptMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&deptMap)
	deptId, cerr := strconv.Atoi(deptMap["deptId"])
	if cerr != nil {
		response.Fail(ctx, nil, cerr.Error())
		return
	}
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

func SendDepartment(ctx *gin.Context) {
	deptId, err1 := strconv.Atoi(ctx.Param("deptId"))
	if err1 != nil {
		response.Fail(ctx, nil, err1.Error())
		return
	}
	department, err := service.GetDepartmentById(int64(deptId))
	if util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, gin.H{
		"department": department,
	}, "获取部门成功！")
}

func UpdateDepartment(ctx *gin.Context) {
	var deptMap = make(map[string]string)
	var department model.Department
	var err *util.Err
	json.NewDecoder(ctx.Request.Body).Decode(&deptMap)
	deptId, err1 := strconv.Atoi(deptMap["deptId"])
	if err1 != nil {
		response.Fail(ctx, nil, err1.Error())
		return
	}
	department.DeptId = int64(deptId)
	department.DeptName = deptMap["deptName"]
	department.Location = deptMap["location"]
	department.WorkTime = deptMap["workTime"]
	department.Description = deptMap["description"]
	department.Picture = deptMap["picture"]
	department.Contact = deptMap["contact"]
	if err = service.UpdateDepartment(&department); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "部门信息更新成功！")
}
