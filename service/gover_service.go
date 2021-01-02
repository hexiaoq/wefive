package service

import (
	"gover-server/common"
	"gover-server/model"
	"gover-server/util"
	"log"
)

func AddGovernor(phone string, deptName string) *util.Err {
	var department *model.Department
	var err *util.Err
	if department, err = GetDepartmentByName(deptName); util.IsFailed(err) {
		log.Println(err.Message)
		return err
	}
	var governor model.Governor
	governor.Phone = phone
	governor.DeptId = department.DeptId
	governor.Password = phone
	db := common.GetDB()
	if err1 := db.Create(&governor).Error; err1 != nil {
		log.Println(err1)
		return util.Fail(err1.Error())
	}
	return util.Success()
}

func DeleteGovernorByPhone(phone string) *util.Err {
	db := common.GetDB()
	if err := db.Where("phone = ?", phone).Delete(&model.Governor{}).Error; err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}
	return util.Success()
}

func DeleteGovernorByDeptId(deptId int64) *util.Err {
	db := common.GetDB()
	err := db.Where("dept_id = ?", deptId).Delete(&model.Governor{}).Error
	if err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}
	return util.Success()
}

func GetAllGovernors() (*[]model.GovernorDto, *util.Err) {
	db := common.GetDB()
	var governors []model.Governor
	var governorDtos []model.GovernorDto
	err := db.Find(&governors).Error
	if err != nil {
		log.Println(err.Error())
		return nil, util.Fail(err.Error())
	}
	for _, gover := range governors {
		goverDto, err1 := ToGovernorDto(&gover)
		if util.IsFailed(err1) {
			log.Println(err1.Message)
			continue
		}
		governorDtos = append(governorDtos, *goverDto)
	}
	return &governorDtos, util.Success()
}

func ToGovernorDto(governor *model.Governor) (*model.GovernorDto, *util.Err) {
	department, err := GetDepartmentById(governor.DeptId)
	if util.IsFailed(err) {
		log.Println(err)
		return nil, util.Fail(err.Message)
	}
	var governorDto model.GovernorDto
	governorDto.DeptId = governor.DeptId
	governorDto.DeptName = department.DeptName
	governorDto.Phone = governor.Phone
	governorDto.GoverId = governor.GoverId
	return &governorDto, util.Success()
}
