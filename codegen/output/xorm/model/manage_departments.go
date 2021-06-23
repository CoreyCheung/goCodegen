package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.978334 +0800 CST m=+0.090347912
type ManageDepartments struct {
    CreatedAt        date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt        *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    DepartmentLevel1 string         `xorm:"'department_level1' VARCHAR(255)" json:"departmentLevel1,omitempty"`
    DepartmentLevel2 string         `xorm:"'department_level2' VARCHAR(255)" json:"departmentLevel2,omitempty"`
    ID               uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    Level            int            `xorm:"'level' INT" json:"level"`
    Remark           string         `xorm:"'remark' VARCHAR(255)" json:"remark,omitempty"`
    Reserved         string         `xorm:"'reserved' VARCHAR(255)" json:"reserved,omitempty"`
    UpdatedAt        date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (ManageDepartments) TableName() string {
	return "manage_departments"
}

func NewManageDepartments() ManageDepartments {
    return ManageDepartments{}
}
