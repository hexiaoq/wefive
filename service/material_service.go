package service

import (
	"log"
	"wefive/common"
	"wefive/model"
	"wefive/util"
)

// 根据业务id获取所有材料
func GetMaterialsByBusId(busId int64) (*[]model.Material, *util.Err) {
	db := common.GetDB()
	var materials []model.Material
	err := db.Where("bus_id = ?", busId).Find(&materials).Error
	if err != nil {
		log.Println(err.Error())
		return nil, util.Fail(err.Error())
	}
	return &materials, util.Success()
}

func AddMaterial(material *model.Material, busId int64) (int64, *util.Err) {
	db := common.GetDB()
	material.BusId = busId
	err := db.Create(material).Error
	if err != nil {
		log.Println(err)
		return 0, util.Fail(err.Error())
	}
	return material.MaterialId, util.Success()
}

func UpdateMaterial(material *model.Material) *util.Err {
	db := common.GetDB()
	err := db.Save(material).Error
	if err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}
	return util.Success()
}

func DeleteMaterial(materialId int64) *util.Err {
	db := common.GetDB()
	if err := db.Where("material_id = ?", materialId).Delete(&model.Material{}).Error; err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}
	// 在流程材料表中删除
	if err := db.Where("material_id = ?", materialId).Delete(&model.ProcessMaterial{}).Error; err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}
	return util.Success()
}

func DeleteMaterialByBusId(busId int64) *util.Err {
	db := common.GetDB()
	err := db.Where("bus_id = ?", busId).Delete(&model.Material{}).Error
	if err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}
	return util.Success()
}

// 根据流程id获取所有材料
func GetMaterialsByProcessId(processId int64) (*[]model.Material, *util.Err) {
	db := common.GetDB()
	var materials []model.Material

	// 查流程材料表
	var processMaterials []model.ProcessMaterial
	if err := db.Where("process_id = ?", processId).Find(&processMaterials).Error; err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}

	// 按材料id获取材料
	for _, pm := range processMaterials {
		material, err1 := GetMaterialById(pm.MaterialId)
		if util.IsFailed(err1) {
			log.Println(err1)
			continue
		}
		materials = append(materials, *material)
	}
	return &materials, util.Success()
}

func GetMaterialById(materialId int64) (*model.Material, *util.Err) {
	db := common.GetDB()
	var material model.Material
	err := db.Where("material_id = ?", materialId).First(&material).Error
	if err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}
	return &material, util.Success()
}
