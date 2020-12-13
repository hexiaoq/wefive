package service

import (
	"log"
	"wefive/common"
	"wefive/model"
	"wefive/util"
)

func GetAllBusinessesByDeptName(deptName string) (*[]model.Business, *util.Err) {
	var businesses []model.Business
	// 获取部门id
	deptId, err1 := GetDeptIdByName(deptName)
	if util.IsFailed(err1) {
		log.Println(err1)
		return nil, err1
	}
	db := common.GetDB()
	err := db.Where("dept_id = ?", deptId).Find(&businesses).Error
	if err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}
	return &businesses, util.Success()
}

func AddBusiness(business *model.Business, deptName string) *util.Err {
	deptId, err1 := GetDeptIdByName(deptName)
	if util.IsFailed(err1) {
		log.Println(err1)
		return err1
	}
	business.DeptId = deptId
	db := common.GetDB()
	// 在business表添加
	err := db.Create(business).Error
	if err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}
	return util.Success()
}

// 根据业务id获取业务
func GetBusinessById(busId int64) (*model.Business, *util.Err) {
	db := common.GetDB()
	var business model.Business
	err := db.Where("bus_id = ?", busId).First(&business).Error
	if err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}
	return &business, util.Success()
}

// 根据业务id删除业务
func DeleteBusiness(busId int64) *util.Err {
	db := common.GetDB()
	// 在business表中删除
	err := db.Where("bus_id = ?", busId).Delete(&model.Business{}).Error
	if err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}

	// 在material表中删除
	if err := DeleteMaterialByBusId(busId); util.IsFailed(err) {
		log.Println(err)
		return util.Fail(err.Message)
	}

	return util.Success()
}

func DeleteBusinessByDeptId(deptId int64) *util.Err {
	db := common.GetDB()
	err := db.Where("dept_id = ?", deptId).Delete(&model.Business{}).Error
	if err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}
	return util.Success()
}

func UpdateBusiness(business *model.Business) *util.Err {
	db := common.GetDB()
	err := db.Save(business).Error
	if err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}
	return util.Success()
}
