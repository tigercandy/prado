package system

import (
	"github.com/gin-gonic/gin"
	"github.com/tigercandy/prado/global/orm"
	"github.com/tigercandy/prado/internal/models/common/response"
	"github.com/tigercandy/prado/internal/models/system"
	"github.com/tigercandy/prado/pkg/utils"
)

func GetSettings(c *gin.Context) {
	var (
		err      error
		classify string
		data     system.Settings
	)
	classify = c.DefaultQuery("classify", "")
	if classify != "" {
		data.Classify, _ = utils.String2Int(classify)
	}

	result, err := data.Get()
	if err != nil {
		response.Error(c, -1, "")
		return
	}

	response.Success(c, result, "")
}

func SetSettings(c *gin.Context) {
	var (
		err           error
		settings      system.Settings
		settingsTotal int64
	)
	err = c.ShouldBind(&settings)
	if err != nil {
		response.Error(c, -1, "")
		return
	}

	err = orm.Eloquent.Model(&system.Settings{}).Where("classify = ? ", settings.Classify).Count(&settingsTotal).Error
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	if settingsTotal == 0 {
		err = orm.Eloquent.Create(&settings).Error
		if err != nil {
			response.Error(c, -1, "")
			return
		}
	} else {
		err = orm.Eloquent.Model(&settings).Where("classify = ? ", settings.Classify).Updates(&settings).Error
		if err != nil {
			response.Error(c, -1, "")
			return
		}
	}

	response.Success(c, "", "配置信息保存成功!")
}
