package model

type Governor struct {
	GoverId      int64  `json:"gover_id,omitempty" gorm:"type:bigint;primary_key;not null;AUTO_INCREMENT"`
	DeptId       int64 `json:"dept_id,omitempty" gorm:"type:bigint;not null"`
	Password 	 string `json:"password,omitempty" gorm:"type:varchar(50);default:null"`
	Phone   	 string `json:"phone,omitempty" gorm:"type:varchar(20);default:null"`
}

type GovernorDto struct {
	GoverId      int64  `json:"gover_id,omitempty" gorm:"type:bigint;primary_key;not null;AUTO_INCREMENT"`
	DeptId       int64 `json:"dept_id,omitempty" gorm:"type:bigint;not null"`
	DeptName     string `json:"dept_name,omitempty" gorm:"type:varchar(20);not null;unique;"`
	Phone   	 string `json:"phone,omitempty" gorm:"type:varchar(20);default:null"`
}
