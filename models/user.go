package models

import "time"

// User 用户
type User struct {
	// Id 主键
	ID int64 `json:"id" form:"id" xorm:"pk id"`
	// Avatar 头像
	Avatar string `json:"avatar" form:"avatar"`
	// Account 账号
	Account string `json:"account" form:"account"`
	// Password 密码
	Password string `json:"password" form:"password"`
	// RePassword
	RePassword string `form:"rePassword" xorm:"-"`
	// Salt md5密码盐
	Salt string `json:"salt"`
	// Name 名称
	Name string `json:"name" form:"name"`
	// Birthday 生日
	Birthday time.Time `json:"birthday"`
	// Sex 性别
	Sex     int8   `json:"sex" form:"sex"`
	SexName string `json:"sexname" xorm:"<- sexname"`
	// Email 电子邮件
	Email string `json:"email" form:"email"`
	// Phone 电话
	Phone string `json:"phone" form:"phone"`
	// RoleId 角色ID
	RoleID string `json:"roleid" xorm:"roleid"`
	// DeptId 部门Id
	DeptID int `json:"deptid" form:"deptid" xorm:"deptid"`
	// Status 状态(1：启用  2：冻结  3：删除）
	Status     int8   `json:"status"`
	StatusName string `json:"statusname" xorm:"<- statusname"`
	// CreateAt 创建时间
	CreateTime time.Time `json:"createTime" xorm:"created 'createtime'"`
}

// UserRole 用户角色
type UserRole struct {
	User `xorm:"extends"`
	Role `xorm:"extends"`
}

// TableName set table
func (User) TableName() string {
	return "sys_user"
}
