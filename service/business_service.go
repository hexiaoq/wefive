package service

import (
	"claps-admin/common"
	"claps-admin/model"
	"claps-admin/util"
	"log"
)

func GetAllBusinessesByDeptName(deptName string) (*[]model.Business, *util.Err) {
	var businesses []model.Business

	// 获取部门id
	deptId, _ := GetDeptIdByName(deptName)

	// 查dept_business表获取busId
	deptBus, err1 := GetAllBusIdsByDeptId(deptId)
	if util.IsFailed(err1) {
		log.Println(err1)
		return nil, util.Fail(err1.Message)
	}
	// 根据获取的busId获取业务
	for _, dpb := range *deptBus {
		var bus *model.Business
		var err2 *util.Err
		if bus, err2 = GetBusinessById(dpb.BusId); util.IsFailed(err2) {
			log.Println(err2.Message)
			continue
		}
		businesses = append(businesses, *bus)
	}
	return &businesses, util.Success()
}

func AddBusiness(business *model.Business, deptName string) *util.Err {
	db := common.GetDB()
	// 在business表添加
	err := db.Create(business).Error
	if err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}
	// 在dept_business表添加
	var deptBus model.DeptBusiness
	var err2 *util.Err
	deptBus.BusId = business.BusId
	deptBus.DeptId, err2 = GetDeptIdByName(deptName)
	if util.IsFailed(err2) {
		log.Println(err2)
		// 失败则删除业务
		DeleteBusinessById(business.BusId)
		return util.Fail(err2.Message)
	}
	err1 := db.Create(&deptBus).Error
	if err1 != nil {
		log.Println(err1)
		// 失败则删除业务
		DeleteBusinessById(business.BusId)
		log.Println("添加失败业务已删除。busId : ", business.BusId)
		return util.Fail(err1.Error())
	}
	return util.Success()
}

// 根据部门id获取多个业务id
func GetAllBusIdsByDeptId(deptId int64) (*[]model.DeptBusiness, *util.Err) {
	var deptBus []model.DeptBusiness
	db := common.GetDB()
	err := db.Where("dept_id = ?", deptId).Find(&deptBus).Error
	if err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}
	return &deptBus, util.Success()
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
func DeleteBusinessById(busId int64) *util.Err {
	db := common.GetDB()
	// 在business表中删除
	err := db.Where("bus_id = ?", busId).Delete(&model.Business{}).Error
	if err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}

	// 在dept_business表中删除
	err1 := db.Where("bus_id = ?", busId).Delete(&model.DeptBusiness{}).Error
	if err1 != nil {
		log.Println(err1)
		return util.Fail(err1.Error())
	}

	// 在material表中删除

	return util.Success()
}
