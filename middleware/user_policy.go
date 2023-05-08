package middleware

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func UserPolicy() *casbin.Enforcer {

	adapter, err := gormadapter.NewAdapter("mysql", "root:@tcp(localhost:3306)/")
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}
	enforcer, err := casbin.NewEnforcer("./config/rbac_model.conf", adapter)
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}
	if hasPolicy := enforcer.HasPolicy("admin", "data", "read"); !hasPolicy {
		enforcer.AddPolicy("admin", "data", "read")
	}
	if hasPolicy := enforcer.HasPolicy("admin", "data", "write"); !hasPolicy {
		enforcer.AddPolicy("admin", "data", "write")
	}
	if hasPolicy := enforcer.HasPolicy("user", "data", "read"); !hasPolicy {
		enforcer.AddPolicy("user", "data", "read")
	}
	return enforcer
}
