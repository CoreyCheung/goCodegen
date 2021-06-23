package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.68607 +0800 CST m=+0.089386221
type ManageRoles struct {
    AccessControl    string         `gorm:"column:access_control;type:varchar(255);size:1020" json:"accessControl,omitempty"`
    CreatedAt        date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt        *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    DepartmentLevel1 string         `gorm:"column:department_level1;type:varchar(255);size:1020" json:"departmentLevel1,omitempty"`
    DepartmentLevel2 string         `gorm:"column:department_level2;type:varchar(255);size:1020" json:"departmentLevel2,omitempty"`
    ID               uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    Remark           string         `gorm:"column:remark;type:varchar(255);size:1020" json:"remark,omitempty"`
    RoleName         string         `gorm:"column:role_name;type:varchar(255);size:1020" json:"roleName,omitempty"`
    UpdatedAt        date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (ManageRoles) TableName() string {
	return "manage_roles"
}
func (this ManageRoles) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this ManageRoles) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this ManageRoles) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this ManageRoles) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(ManageRoles{}).Error; err != nil {
		return err
	}
	return nil
}

func NewManageRoles() ManageRoles {
    return ManageRoles{}
}

