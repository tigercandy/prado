package middleware

import (
	"github.com/tigercandy/prado/global"
	"github.com/tigercandy/prado/internal/handler"
	"github.com/tigercandy/prado/pkg/jwt"
	"time"
)

func InitAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "prado zone",
		Key:             []byte(global.App.Config.Jwt.Secret),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		PayloadFunc:     handler.PayloadFunc,
		IdentityHandler: handler.IdentityHandler,
		Authenticator:   handler.Authenticator,
		Authorizator:    handler.Authorizator,
		Unauthorized:    handler.Unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
}
