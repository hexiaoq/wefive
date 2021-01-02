package model

type Department struct {
	DeptId      int64  `json:"dept_id" gorm:"type:bigint;primary_key;not null;AUTO_INCREMENT"`
	DeptName    string `json:"dept_name" gorm:"type:varchar(20);not null;unique;"`
	Location    string `json:"location" gorm:"type:varchar(50);default:null"`
	WorkTime    string `json:"work_time" gorm:"type:varchar(100);default:null"`
	Description string `json:"description" gorm:"type:varchar(200);default:null"`
	Picture     string `json:"picture"`
	Contact     string `json:"contact"`
	Longtitude  string `json:"longtitude"`
	Latitude    string `json:"latitude"`
}
