package system

import (
	"encoding/json"
	"github.com/tigercandy/prado/global/orm"
	"github.com/tigercandy/prado/internal/models/common"
)

type Settings struct {
	Classify int             `gorm:"column:classify;type:int(11)" json:"classify" form:"classify"`
	Content  json.RawMessage `gorm:"column:content;type:json" json:"content" form:"content"`
	common.Model
}

func (Settings) TableName() string {
	return "sys_settings"
}

func (s *Settings) Get() ([]Settings, error) {
	var doc []Settings
	table := orm.Eloquent.Table(s.TableName())
	if s.Classify != 0 {
		table = table.Where("classify = ? ", s.Classify)
	}

	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}

	return doc, nil
}
