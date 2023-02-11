package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/tigercandy/prado/internal/models/common/request"
	"github.com/tigercandy/prado/internal/models/system"
	ss "github.com/tigercandy/prado/internal/services/system"
	"github.com/tigercandy/prado/pkg/jwt"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(system.SysUser)
		r, _ := v["role"].(system.SysRole)
		return jwt.MapClaims{
			jwt.IdentityKey: u.UserId,
			jwt.RoleIdKey:   r.RoleId,
			jwt.RoleKey:     r.RoleKey,
			jwt.NiceKey:     u.Username,
			jwt.RoleNameKey: r.RoleName,
		}
	}

	return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims["identity"],
		"UserName":    claims["nice"],
		"RoleKey":     claims["roleKey"],
		"UserId":      claims["identity"],
		"RoleIds":     claims["roleId"],
	}
}

func Authenticator(c *gin.Context) (interface{}, error) {
	var (
		err          error
		form         request.Login
		isVerifyCode interface{}
	)

	if err = c.ShouldBind(&form); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	isVerifyCode, err = ss.SysSettingsServer.GetContentByKey("1", "is_verify_code")
	if err != nil {
		return nil, errors.New("获取验证码登录信息失败")
	}

	if isVerifyCode != nil && isVerifyCode.(bool) {
		if !store.Verify(form.UUID, form.Code, true) {
			return nil, jwt.ErrInvalidVerifyCode
		}
	}

	user, role, e := ss.SysUserService.Login(form)

	if e == nil {
		if user.Status == system.StatusForbidden {
			return nil, errors.New("用户被禁用!")
		}

		return map[string]interface{}{"user": user, "role": role}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

func Authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(system.SysUser)
		r, _ := v["role"].(system.SysRole)
		c.Set("userId", u.UserId)
		c.Set("userName", u.UserName)
		c.Set("role", r.RoleName)
		c.Set("roleIds", r.RoleId)

		return true
	}

	return false
}

func Unauthorized(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "退出成功",
	})
}
