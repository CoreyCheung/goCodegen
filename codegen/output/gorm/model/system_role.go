package model

// Created by CodeMachine on 2020-05-10 15:46:04.736711 +0800 CST m=+0.140026445
// [系统] 管理角色
type SystemRole struct {
    // 角色权限
    Auth   string `gorm:"not null;column:auth;type:text;size:65535" json:"auth,omitempty"`
    // 创建时间
    Ctime  uint   `gorm:"not null;column:ctime;type:int unsigned" json:"ctime"`
    ID     uint   `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    // 角色简介
    Intro  string `gorm:"not null;column:intro;type:varchar(200);size:600" json:"intro,omitempty"`
    // 修改时间
    Mtime  uint   `gorm:"not null;column:mtime;type:int unsigned" json:"mtime"`
    // 角色名称
    Name   string `gorm:"not null;column:name;type:varchar(50);size:150" json:"name,omitempty"`
    // 状态
    Status int8   `gorm:"not null;column:status;type:tinyint unsigned" json:"status"`
}

func (SystemRole) TableName() string {
	return "system_role"
}
func (this SystemRole) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemRole) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemRole) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this SystemRole) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(SystemRole{}).Error; err != nil {
		return err
	}
	return nil
}

func NewSystemRole() SystemRole {
    return SystemRole{}
}

