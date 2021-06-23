package model

// Created by CodeMachine on 2020-05-10 15:46:04.720035 +0800 CST m=+0.123350883
// [系统] 钩子表
type SystemHook struct {
    // 创建时间
    Ctime  uint   `gorm:"not null;column:ctime;type:int unsigned" json:"ctime"`
    ID     uint   `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    // 钩子简介
    Intro  string `gorm:"not null;column:intro;type:varchar(200);size:600" json:"intro,omitempty"`
    // 更新时间
    Mtime  uint   `gorm:"not null;column:mtime;type:int unsigned" json:"mtime"`
    // 钩子名称
    Name   string `gorm:"not null;column:name;type:varchar(50);size:150" json:"name,omitempty"`
    // 钩子来源[plugins.插件名，module.模块名]
    Source string `gorm:"not null;column:source;type:varchar(50);size:150" json:"source,omitempty"`
    Status int8   `gorm:"not null;column:status;type:tinyint unsigned" json:"status"`
    // 系统插件
    System int8   `gorm:"not null;column:system;type:tinyint unsigned" json:"system"`
}

func (SystemHook) TableName() string {
	return "system_hook"
}
func (this SystemHook) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemHook) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemHook) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this SystemHook) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(SystemHook{}).Error; err != nil {
		return err
	}
	return nil
}

func NewSystemHook() SystemHook {
    return SystemHook{}
}

