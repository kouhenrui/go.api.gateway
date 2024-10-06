package config

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
)

type CasbinEnforcer struct {
	enforcer *casbin.Enforcer
}

// NewCasbinEnforcer 初始化 Casbin 并连接 MySQL 数据库
func NewCasbinEnforcer(dsn string) (*CasbinEnforcer, error) {
	// 创建 MySQL 数据库连接
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	return nil, err
	//}
	Type := "mysql"
	db := "" //CabinConfig.UserName + ":" + CabinConfig.PassWord + "@tcp(" + CabinConfig.HOST + ":" + CabinConfig.Port + ")/"
	// 初始化 Gorm Adapter
	adapter, err := gormadapter.NewAdapter(Type, db, true)
	if err != nil {
		return nil, err
	}

	// 初始化 Casbin enforcer
	e, err := casbin.NewEnforcer("auth/model.conf", adapter)
	if err != nil {
		return nil, err
	}

	// 加载策略
	if err = e.LoadPolicy(); err != nil {
		return nil, err
	}

	return &CasbinEnforcer{enforcer: e}, nil
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
	ce.enforcer.SavePolicy() // 保存到数据库
	return nil
}

// RemovePolicy 删除权限策略
func (ce *CasbinEnforcer) RemovePolicy(sub, obj, act string) error {
	ok, err := ce.enforcer.RemovePolicy(sub, obj, act)
	if !ok {
		return err
	}
	ce.enforcer.SavePolicy() // 保存更改
	return nil
}

// AddRoleForUser 添加角色
func (ce *CasbinEnforcer) AddRoleForUser(user, role string) error {
	ok, err := ce.enforcer.AddRoleForUser(user, role)
	if !ok {
		return err
	}
	ce.enforcer.SavePolicy()
	return nil
}

// DeleteRoleForUser 删除用户的角色
func (ce *CasbinEnforcer) DeleteRoleForUser(user, role string) error {
	ok, err := ce.enforcer.DeleteRoleForUser(user, role)
	if !ok {
		return err
	}
	ce.enforcer.SavePolicy()
	return nil
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
	ce.enforcer.SavePolicy()
	return nil
}

// GetAllPolicies 获取所有的策略
func (ce *CasbinEnforcer) GetAllPolicies() ([][]string, error) {
	return ce.enforcer.GetPolicy()
}
