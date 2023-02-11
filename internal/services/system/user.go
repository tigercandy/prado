package system

import (
	"github.com/gin-gonic/gin"
	"github.com/tigercandy/prado/global/orm"
	"github.com/tigercandy/prado/internal/models/common/request"
	"github.com/tigercandy/prado/internal/models/system"
	"github.com/tigercandy/prado/pkg/jwt"
	"github.com/tigercandy/prado/pkg/utils"
)

type sysUserService struct {
}

var SysUserService = new(sysUserService)

func (sus *sysUserService) Login(params request.Login) (user system.SysUser, role system.SysRole, err error) {
	var (
		userModel system.SysUser
		roleModel system.SysRole
	)
	err = orm.Eloquent.Table(userModel.TableName()).Where("username = ? ", params.Username).First(&user).Error
	if err != nil {
		return
	}

	_, err = utils.BcryptMakeCheck([]byte(params.Password), user.Password)
	if err != nil {
		return
	}

	err = orm.Eloquent.Table(roleModel.TableName()).Where("role_id = ? ", user.RoleId).First(&role).Error
	if err != nil {
		return
	}

	return
}

func ExtractClaims(c *gin.Context) jwt.MapClaims {
	claims, exists := c.Get("JWT_PAYLOAD")
	if !exists {
		return make(jwt.MapClaims)
	}

	return claims.(jwt.MapClaims)
}

func (sus *sysUserService) GetRoleName(c *gin.Context) string {
	data := ExtractClaims(c)
	if data["roleKey"] != nil {
		return (data["roleKey"]).(string)
	}

	return ""
}

func (sus *sysUserService) GetRoleId(c *gin.Context) int {
	data := ExtractClaims(c)
	if data["roleId"] != nil {
		i := int((data["roleId"]).(float64))
		return i
	}

	return 0
}

func (sus *sysUserService) GetUserId(c *gin.Context) int {
	data := ExtractClaims(c)
	if data["identity"] != nil {
		return int((data["identity"]).(float64))
	}

	return 0
}

func (sus *sysUserService) GetUserIdStr(c *gin.Context) string {
	data := ExtractClaims(c)
	if (data["identity"]) != nil {
		return utils.Int642String(int64(data["identity"].(float64)))
	}
	return ""
}
