package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.691025 +0800 CST m=+0.094341419
type ManageUsers struct {
    AccountId       string         `gorm:"column:account_id;type:varchar(255);size:1020" json:"accountId,omitempty"`
    CreatedAt       date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt       *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    ID              uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    Mail            string         `gorm:"column:mail;type:varchar(255);size:1020" json:"mail,omitempty"`
    Mobile          string         `gorm:"column:mobile;type:varchar(255);size:1020" json:"mobile,omitempty"`
    Name            string         `gorm:"column:name;type:varchar(255);size:1020" json:"name,omitempty"`
    Password        string         `gorm:"column:password;type:varchar(255);size:1020" json:"password,omitempty"`
    RoleId          int            `gorm:"column:role_id;type:int" json:"roleId"`
    RoleName        string         `gorm:"column:role_name;type:varchar(255);size:1020" json:"roleName,omitempty"`
    Token           string         `gorm:"column:token;type:varchar(255);size:1020" json:"token,omitempty"`
    UpdateAccountId string         `gorm:"column:update_account_id;type:varchar(255);size:1020" json:"updateAccountId,omitempty"`
    UpdatedAt       date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (ManageUsers) TableName() string {
	return "manage_users"
}
func (this ManageUsers) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this ManageUsers) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this ManageUsers) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this ManageUsers) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(ManageUsers{}).Error; err != nil {
		return err
	}
	return nil
}

func NewManageUsers() ManageUsers {
    return ManageUsers{}
}

