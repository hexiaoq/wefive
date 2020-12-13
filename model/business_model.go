package model

type Business struct {
	BusId       int64   `json:"bus_id,omitempty" gorm:"type:bigint;primary_key;not null;AUTO_INCREMENT"`
	DeptId      int64   `json:"dept_id,omitempty" gorm:"type:bigint;not null;"`
	BusName     string  `json:"bus_name,omitempty" gorm:"type:varchar(50);not null;"`
	Description string  `json:"description,omitempty" gorm:"type:varchar(500);default:null"`
	Requirement string  `json:"requirement,omitempty" gorm:"type:varchar(500);default:null"`
	Cost        float64 `json:"cost,omitempty" gorm:"type:real;default:null"`
}
