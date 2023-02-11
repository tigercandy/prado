package casbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/tigercandy/prado/global"
	"github.com/tigercandy/prado/global/orm"
)

func Casbin() (*casbin.Enforcer, error) {
	conn := orm.MysqlConn
	Apter, err := gormadapter.NewAdapter(global.App.Config.Database.Driver, conn, true)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer("configs/rbac_model.conf", Apter)
	if err != nil {
		return nil, err
	}
	if err := e.LoadPolicy(); err == nil {
		return e, err
	} else {
		return nil, err
	}
}
