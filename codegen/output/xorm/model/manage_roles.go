package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.979173 +0800 CST m=+0.091186673
type ManageRoles struct {
    AccessControl    string         `xorm:"'access_control' VARCHAR(255)" json:"accessControl,omitempty"`
    CreatedAt        date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt        *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    DepartmentLevel1 string         `xorm:"'department_level1' VARCHAR(255)" json:"departmentLevel1,omitempty"`
    DepartmentLevel2 string         `xorm:"'department_level2' VARCHAR(255)" json:"departmentLevel2,omitempty"`
    ID               uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    Remark           string         `xorm:"'remark' VARCHAR(255)" json:"remark,omitempty"`
    RoleName         string         `xorm:"'role_name' VARCHAR(255)" json:"roleName,omitempty"`
    UpdatedAt        date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (ManageRoles) TableName() string {
	return "manage_roles"
}

func NewManageRoles() ManageRoles {
    return ManageRoles{}
}
