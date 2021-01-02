package service

import (
	"gover-server/common"
	"gover-server/model"
	"gover-server/util"
	"log"
)

// 添加一个部门
func AddDepartment(department *model.Department) *util.Err {
	db := common.GetDB()
	if len(department.DeptName) == 0 {
		log.Println("部门名字不能为空！")
		return util.Fail("部门名字不能为空！")
	}
	if IsDepartmentExistByName(department.DeptName) {
		log.Println("部门名字已经存在: ", department.DeptName)
		return util.Fail("部门名字已经存在!")
	}
	err := db.Create(&department).Error
	if err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}
	return util.Success()
}

// 获取所有部门
func GetAllDepartments() (*[]model.Department, *util.Err) {
	db := common.GetDB()
	var departments []model.Department
	err := db.Find(&departments).Error
	if err != nil {
		log.Println(err.Error())
		return nil, util.Fail(err.Error())
	}
	return &departments, util.Success()
}

// 删除部门
func DeleteDepartmentById(deptId int64) *util.Err {
	if !IsDepartmentExistById(deptId) {
		log.Println("要删除的部门不存在！deptId: ", deptId)
		return util.Success()
	}
	db := common.GetDB()
	// 删除部门
	db.Where("dept_id = ?", deptId).Delete(&model.Department{})
	// 删除业务
	if err := DeleteBusinessByDeptId(deptId); util.IsFailed(err) {
		log.Println(err)
		return util.Fail(err.Message)
	}

	// 删除人员
	if err := DeleteGovernorByDeptId(deptId); util.IsFailed(err) {
		log.Println(err)
		return util.Fail(err.Message)
	}

	return util.Success()
}

// 判断部门是否存在
func IsDepartmentExistByName(deptName string) bool {
	db := common.GetDB()
	var dept model.Department
	db.Where("dept_name = ?", deptName).First(&dept)
	return dept.DeptId != 0
}

// 判断部门是否存在
func IsDepartmentExistById(deptId int64) bool {
	db := common.GetDB()
	var dept model.Department
	db.Where("dept_id = ?", deptId).First(&dept)
	return dept.DeptId != 0
}

// 根据id获取部门
func GetDepartmentById(deptId int64) (*model.Department, *util.Err) {
	db := common.GetDB()
	var dept model.Department
	err := db.Where("dept_id = ?", deptId).First(&dept).Error
	if err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}
	return &dept, util.Success()
}

// 根据名字获取部门
func GetDepartmentByName(deptName string) (*model.Department, *util.Err) {
	db := common.GetDB()
	var dept model.Department
	err := db.Where("dept_name = ?", deptName).First(&dept).Error
	if err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}
	return &dept, util.Success()
}

// 获取部门id
func GetDeptIdByName(deptName string) (int64, *util.Err) {
	dept, err := GetDepartmentByName(deptName)
	if util.IsFailed(err) {
		log.Println(err.Message)
		return 0, util.Fail(err.Message)
	}
	return dept.DeptId, util.Success()
}

func UpdateDepartment(department *model.Department) *util.Err {
	depart, err := GetDepartmentById(department.DeptId)
	if util.IsFailed(err) {
		log.Println(err)
		return err
	}
	if len(department.DeptName) > 0 {
		depart.DeptName = department.DeptName
	}
	if len(department.Contact) > 0 {
		depart.Contact = department.Contact
	}
	if len(department.Picture) > 0 {
		depart.Picture = department.Picture
	}
	if len(department.Location) > 0 {
		depart.Location = department.Location
	}
	if len(department.Description) > 0 {
		depart.Description = department.Description
	}
	if len(department.WorkTime) > 0 {
		depart.WorkTime = department.WorkTime
	}
	db := common.GetDB()
	err1 := db.Save(depart).Error
	if err1 != nil {
		log.Println(err1.Error())
		return util.Fail(err1.Error())
	}
	return util.Success()
}
