package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.983183 +0800 CST m=+0.095197250
type ManageUsers struct {
    AccountId       string         `xorm:"'account_id' VARCHAR(255)" json:"accountId,omitempty"`
    CreatedAt       date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt       *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    ID              uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    Mail            string         `xorm:"'mail' VARCHAR(255)" json:"mail,omitempty"`
    Mobile          string         `xorm:"'mobile' VARCHAR(255)" json:"mobile,omitempty"`
    Name            string         `xorm:"'name' VARCHAR(255)" json:"name,omitempty"`
    Password        string         `xorm:"'password' VARCHAR(255)" json:"password,omitempty"`
    RoleId          int            `xorm:"'role_id' INT" json:"roleId"`
    RoleName        string         `xorm:"'role_name' VARCHAR(255)" json:"roleName,omitempty"`
    Token           string         `xorm:"'token' VARCHAR(255)" json:"token,omitempty"`
    UpdateAccountId string         `xorm:"'update_account_id' VARCHAR(255)" json:"updateAccountId,omitempty"`
    UpdatedAt       date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (ManageUsers) TableName() string {
	return "manage_users"
}

func NewManageUsers() ManageUsers {
    return ManageUsers{}
}
