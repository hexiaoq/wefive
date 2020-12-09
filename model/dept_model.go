package model

type Department struct {
	DeptId      int64  `json:"dept_id,omitempty" gorm:"type:bigint;primary_key;not null;AUTO_INCREMENT"`
	DeptName       string `json:"dept_name,omitempty" gorm:"type:varchar(20);not null;unique;"`
	Location 	 string `json:"location,omitempty" gorm:"type:varchar(50);default:null"`
	WorkTime   	 string `json:"word_time,omitempty" gorm:"type:varchar(100);default:null"`
	Description  string `json:"description,omitempty" gorm:"type:varchar(200);default:null"`
}
