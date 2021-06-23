package model

// Created by CodeMachine on 2020-05-09 22:05:35.050171 +0800 CST m=+0.162182924
// [系统] 管理角色
type SystemRole struct {
    // 角色权限
    Auth   string `xorm:"'auth' not null TEXT" json:"auth,omitempty"`
    // 创建时间
    Ctime  uint   `xorm:"'ctime' not null INT UNSIGNED" json:"ctime"`
    ID     uint   `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    // 角色简介
    Intro  string `xorm:"'intro' not null VARCHAR(200)" json:"intro,omitempty"`
    // 修改时间
    Mtime  uint   `xorm:"'mtime' not null INT UNSIGNED" json:"mtime"`
    // 角色名称
    Name   string `xorm:"'name' not null VARCHAR(50)" json:"name,omitempty"`
    // 状态
    Status int8   `xorm:"'status' not null TINYINT UNSIGNED" json:"status"`
}

func (SystemRole) TableName() string {
	return "system_role"
}

func NewSystemRole() SystemRole {
    return SystemRole{}
}
