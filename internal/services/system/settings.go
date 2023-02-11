package system

import (
	"encoding/json"
	"github.com/tigercandy/prado/global/orm"
	"github.com/tigercandy/prado/internal/models/system"
)

type sysSettingsService struct {
}

var SysSettingsServer = new(sysSettingsService)

func (sss *sysSettingsService) GetContent(classify string) (content map[string]interface{}, err error) {
	var settings system.Settings

	if err = orm.Eloquent.Where("classify = ? ", classify).Find(&settings).Error; err != nil {
		return
	}

	if err = json.Unmarshal(settings.Content, &content); err != nil {
		return
	}

	return
}

func (sss *sysSettingsService) GetContentByKey(classify, key string) (value interface{}, err error) {
	var content map[string]interface{}

	content, err = sss.GetContent(classify)
	if err != nil {
		return
	}

	value = content[key]
	return
}
