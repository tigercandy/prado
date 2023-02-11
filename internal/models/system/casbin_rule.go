package system

type CasbinRule struct {
	Id    int    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	PType string `gorm:"type:varchar(128);" json:"p_type"`
	V0    string `gorm:"type:varchar(128);" json:"v0"`
	V1    string `gorm:"type:varchar(128);" json:"v1"`
	V2    string `gorm:"type:varchar(128);" json:"v2"`
	V3    string `gorm:"type:varchar(128);" json:"v3"`
	V4    string `gorm:"type:varchar(128);" json:"v4"`
	V5    string `gorm:"type:varchar(128);" json:"v5"`
	BaseModel
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}
