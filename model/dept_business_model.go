package model

type DeptBusiness struct {
	DeptId      int64  `json:"dept_id,omitempty" gorm:"type:bigint;primary_key;not null;"`
	BusId      	int64  `json:"bus_id,omitempty" gorm:"type:bigint;primary_key;not null;"`
}
