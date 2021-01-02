package service

import (
	"gover-server/common"
	"gover-server/model"
	"gover-server/util"
	"log"
)

func GetAllBusinessesByDeptName(deptName string) (*[]model.Business, *util.Err) {
	// 获取部门id
	deptId, err1 := GetDeptIdByName(deptName)
	if util.IsFailed(err1) {
		log.Println(err1)
		return nil, err1
	}
	businesses, err := GetAllBusinessesByDeptId(deptId)
	if util.IsFailed(err) {
		log.Println(err)
		return nil, err
	}
	return businesses, util.Success()
}

func GetAllBusinesses() (*[]model.Business, *util.Err) {
	var businesses []model.Business
	db := common.GetDB()
	err := db.Find(&businesses).Error
	if err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}
	return &businesses, util.Success()
}

func GetAllBusinessesByDeptId(deptId int64) (*[]model.Business, *util.Err) {
	var businesses []model.Business
	db := common.GetDB()
	err := db.Where("dept_id = ?", deptId).Find(&businesses).Error
	if err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}
	return &businesses, util.Success()
}

// 获取热门业务
func GetHotBusiness() (*[]model.BusinessDto, *util.Err) {
	// 还没实现，先获取所有业务
	var businesses []model.Business
	var businessesDto []model.BusinessDto
	db := common.GetDB()
	err := db.Find(&businesses).Error
	if err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}

	// 转化dto
	for _, bus := range businesses {
		busDto, err1 := toBusinessDto(&bus)
		if util.IsFailed(err1) {
			log.Println(err1)
			continue
		}
		businessesDto = append(businessesDto, *busDto)
	}
	return &businessesDto, util.Success()
}

func AddBusiness(business *model.Business, deptId int64) *util.Err {
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

	// 在流程表中删除
	if err := DeleteProcessByBusId(busId); util.IsFailed(err) {
		log.Println(err)
		return err
	}

	// 在material表中删除
	if err := DeleteMaterialByBusId(busId); util.IsFailed(err) {
		log.Println(err)
		return util.Fail(err.Message)
	}

	// 在business表中删除
	err := db.Where("bus_id = ?", busId).Delete(&model.Business{}).Error
	if err != nil {
		log.Println(err)
		return util.Fail(err.Error())
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

// 转化业务数据传输类型
func toBusinessDto(business *model.Business) (*model.BusinessDto, *util.Err) {
	var businessDto model.BusinessDto
	deptId := business.DeptId
	department, err := GetDepartmentById(deptId)
	if util.IsFailed(err) {
		log.Println(err)
		return nil, err
	}

	businessDto.BusId = business.BusId
	businessDto.BusName = business.BusName
	businessDto.Requirement = business.Requirement
	businessDto.Description = business.Description
	businessDto.Cost = business.Cost

	businessDto.DeptId = deptId
	businessDto.Phone = department.Contact
	businessDto.Location = department.Location
	businessDto.DeptName = department.DeptName
	businessDto.WorkTime = department.WorkTime
	businessDto.Picture = department.Picture
	return &businessDto, util.Success()
}

// 添加业务模板
func AddBusTemplate(busName string, deptId int64) *util.Err {
	db := common.GetDB()
	var oldBus model.Business
	var newBus model.Business
	// 在业务表中添加
	db.Where("bus_name = ?", busName).First(&oldBus)
	newBus.DeptId = deptId
	newBus.BusName = busName
	newBus.Description = oldBus.Description
	newBus.Cost = oldBus.Cost
	newBus.Requirement = oldBus.Requirement
	db.Create(&newBus)
	oldBusId := oldBus.BusId
	newBusId := newBus.BusId

	// 在材料表中添加
	materials, err := GetMaterialsByBusId(oldBusId)
	if util.IsFailed(err) {
		log.Println(err)
		return err
	}
	for _, m := range *materials {
		var material model.Material
		material.BusId = newBusId
		material.Description = m.Description
		material.PhotoUrl = m.PhotoUrl
		material.MaterialName = m.MaterialName
		db.Create(&material)
	}

	// 在流程表中添加
	processes, err := GetProcessesByBusId(oldBusId)
	if util.IsFailed(err) {
		log.Println(err)
		return err
	}
	for _, p := range *processes {
		var process model.Process
		process.BusId = newBusId
		process.Step = p.Step
		process.Title = p.Title
		process.Content = p.Content
		db.Create(&process)
	}

	return util.Success()
}
