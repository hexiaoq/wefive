package service

import (
	"fmt"
	"gover-server/common"
	"gover-server/model"
	"gover-server/util"
	"log"
)

func AddProcessForBus(process *model.Process) *util.Err {
	db := common.GetDB()
	err := db.Create(process).Error
	if err != nil {
		log.Println(err.Error())
		return util.Fail(err.Error())
	}
	return util.Success()
}

// 为流程添加多个材料
func AddProcessMaterials(processId int64, materials interface{}) *util.Err {
	db := common.GetDB()
	for _, materialId := range materials.([]interface{}) {
		mId := materialId.(float64)
		var processMaterial model.ProcessMaterial
		processMaterial.MaterialId = int64(mId)
		processMaterial.ProcessId = processId
		if err1 := db.Create(&processMaterial).Error; err1 != nil {
			log.Println(err1)
			return util.Fail(err1.Error())
		}
	}
	return util.Success()
}

func GetProcessesByBusId(busId int64) (*[]model.Process, *util.Err) {
	db := common.GetDB()
	var processes []model.Process
	err := db.Where("bus_id = ?", busId).Find(&processes).Error
	processes = sort(processes)
	if err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}
	return &processes, util.Success()
}

func GetAllProcessOfBus(busId int64) (*[]model.ProcessDto, *util.Err) {
	var processDtos []model.ProcessDto
	processes, err := GetProcessesByBusId(busId)
	if util.IsFailed(err) {
		log.Println(err)
		return nil, err
	}

	for _, process := range *processes {
		var processDto *model.ProcessDto
		var err *util.Err
		if processDto, err = toProcessDto(process); util.IsFailed(err) {
			log.Println(err)
			continue
		}
		processDtos = append(processDtos, *processDto)
	}
	return &processDtos, util.Success()
}

// 将一个流程转为包含材料的dto
func toProcessDto(process model.Process) (*model.ProcessDto, *util.Err) {
	// 获取该流程的所有材料
	materials, err := GetMaterialsByProcessId(process.ProcessId)
	if util.IsFailed(err) {
		log.Println(err)
		return nil, util.Fail(err.Message)
	}
	var processDto model.ProcessDto
	processDto.Materials = *materials
	processDto.ProcessId = process.ProcessId
	processDto.BusId = process.BusId
	processDto.Content = process.Content
	processDto.Title = process.Title
	processDto.Step = process.Step
	return &processDto, util.Success()
}

func DeleteProcessByBusId(busId int64) *util.Err {
	db := common.GetDB()
	// 在process_material表中删除
	var processes []model.Process
	err := db.Where("bus_id = ?", busId).Find(&processes).Error
	if err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}
	for _, process := range processes {
		err = db.Where("process_id = ?", process.ProcessId).Delete(&model.ProcessMaterial{}).Error
		if err != nil {
			log.Println(err)
		}
	}

	// 在process表中删除
	err = db.Where("bus_id = ?", busId).Delete(&model.Process{}).Error
	if err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}
	return util.Success()
}

func DeleteProcessMaterial(processId int64, materialId int64) *util.Err {
	db := common.GetDB()
	fmt.Println(processId, materialId)
	err := db.Where("process_id = ? AND material_id = ?", processId, materialId).Delete(&model.ProcessMaterial{}).Error
	if err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}
	return util.Success()
}

func sort(processes []model.Process) []model.Process {
	if len(processes) <= 1 {
		return processes
	}

	var pro []model.Process
	for i := 0; i < len(processes); i++ {
		var min = 999
		minIndex := 0
		for j := 0; j < len(processes); j++ {
			if min > processes[j].Step {
				minIndex = j
				min = processes[j].Step
			}
		}
		pro = append(pro, processes[minIndex])
		processes[minIndex].Step = 999
	}
	return pro
}
