package config

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql" // MySQL 驱动
	"strings"
)

type CasbinEnforcer struct {
	enforcer *casbin.Enforcer
}

// NewCasbinEnforcer 初始化 Casbin 并连接 MySQL 数据库
func NewCasbinEnforcer(CasbinConfig CasbinConf) error {
	//Type := "mysql"
	db := CasbinConfig.UserName + ":" + CasbinConfig.PassWord + "@tcp(" + CasbinConfig.Host + ":" + CasbinConfig.Port + ")/"
	// 初始化 Gorm Adapter
	adapter, err := gormadapter.NewAdapter(CasbinConfig.Type, db, false)
	if err != nil {
		return err
	}

	// 初始化 Casbin enforcer
	e, err := casbin.NewEnforcer("./src/model.conf", adapter)
	if err != nil {
		return err
	}
	//挂载基础策略
	initCasbin(e)
	// 加载策略
	if err = e.LoadPolicy(); err != nil {
		return err
	}
	_ = &CasbinEnforcer{enforcer: e}
	return nil
}

func initCasbin(e *casbin.Enforcer) {

	lp := []string{
		"p, *, /api/v1/login, *",
		"p, *, /api/v1/info, *",
		"p, *, /api/v1/captcha, Get",
		"p, user, /api/v1/resource, GET",
	}
	for _, s := range lp {
		st := strings.Split(s, ", ")
		v0 := st[1]
		v1 := st[2]
		v2 := st[3]

		_, _ = e.AddPolicy(v0, v1, v2)
	}
	lg := []string{
		"g, alice, admin",
		"g, bob, user",
	}
	for _, t := range lg {
		st := strings.Split(t, ", ")
		user := st[1]
		role := st[2]
		_, _ = e.AddRoleForUser(user, role)
		//e.AddRoleForUser(t)
	}
}

// CheckPermission 检查权限
func (ce *CasbinEnforcer) CheckPermission(sub, obj, act string) (bool, error) {
	ok, err := ce.enforcer.Enforce(sub, obj, act)
	return ok, err
}

// AddPolicy 动态添加权限策略
func (ce *CasbinEnforcer) AddPolicy(sub, obj, act string) error {
	ok, err := ce.enforcer.AddPolicy(sub, obj, act)
	if !ok {
		return err
	}
	return ce.enforcer.SavePolicy() // 保存到数据库
}

// RemovePolicy 删除权限策略
func (ce *CasbinEnforcer) RemovePolicy(sub, obj, act string) error {
	ok, err := ce.enforcer.RemovePolicy(sub, obj, act)
	if !ok {
		return err
	}
	return ce.enforcer.SavePolicy() // 保存更改
}

// AddRoleForUser 添加角色
func (ce *CasbinEnforcer) AddRoleForUser(user, role string) error {
	ok, err := ce.enforcer.AddRoleForUser(user, role)
	if !ok {
		return err
	}
	return ce.enforcer.SavePolicy()
}

// DeleteRoleForUser 删除用户的角色
func (ce *CasbinEnforcer) DeleteRoleForUser(user, role string) error {
	ok, err := ce.enforcer.DeleteRoleForUser(user, role)
	if !ok {
		return err
	}
	return ce.enforcer.SavePolicy()
}

// GetRolesForUser 获取用户的角色
func (ce *CasbinEnforcer) GetRolesForUser(user string) ([]string, error) {
	roles, err := ce.enforcer.GetRolesForUser(user)
	return roles, err
}

// HasRoleForUser 检查用户是否具备某角色
func (ce *CasbinEnforcer) HasRoleForUser(user, role string) (bool, error) {
	ok, err := ce.enforcer.HasRoleForUser(user, role)
	return ok, err
}

// GetPermissionsForRole 获取某角色的权限
func (ce *CasbinEnforcer) GetPermissionsForRole(role string) ([][]string, error) {
	return ce.enforcer.GetPermissionsForUser(role)
}

// DeleteRole 删除角色
func (ce *CasbinEnforcer) DeleteRole(role string) error {
	ok, err := ce.enforcer.DeleteRole(role)
	if !ok {
		return err
	}
	return ce.enforcer.SavePolicy()
}

// GetAllPolicies 获取所有的策略
func (ce *CasbinEnforcer) GetAllPolicies() ([][]string, error) {
	return ce.enforcer.GetPolicy()
}
