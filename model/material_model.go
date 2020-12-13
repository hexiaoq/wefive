package model

type Material struct {
	MaterialId   int64  `json:"material_id,omitempty" gorm:"type:bigint;primary_key;not null;AUTO_INCREMENT"`
	BusId        int64  `json:"bus_id,omitempty" gorm:"type:bigint;not null;"`
	MaterialName string `json:"material_name,omitempty" gorm:"type:varchar(20);not null;"`
	PhotoUrl     string `json:"photo_url,omitempty" gorm:"type:varchar(100);default:null"`
	Description  string `json:"description,omitempty" gorm:"type:varchar(500);default:null"`
}
