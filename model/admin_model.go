package model

type Admin struct {
	AdminId  int64  `json:"admin_id,omitempty" gorm:"type:bigint;primary_key;not null;AUTO_INCREMENT"`
	Name     string `json:"name,omitempty" gorm:"type:varchar(20);unique_index:name_UNIQUE;not null"`
	Password string `json:"password,omitempty" gorm:"type:varchar(50);default:null"`
	Phone    string `json:"phone,omitempty" gorm:"type:varchar(20);default:null"`
}
