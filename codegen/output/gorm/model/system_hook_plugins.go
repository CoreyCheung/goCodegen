package model

// Created by CodeMachine on 2020-05-10 15:46:04.721719 +0800 CST m=+0.125034205
// [系统] 钩子-插件对应表
type SystemHookPlugins struct {
    Ctime   uint   `gorm:"not null;column:ctime;type:int unsigned" json:"ctime"`
    // 钩子id
    Hook    string `gorm:"not null;column:hook;type:varchar(32);size:96" json:"hook,omitempty"`
    ID      uint   `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    Mtime   uint   `gorm:"not null;column:mtime;type:int unsigned" json:"mtime"`
    // 插件标识
    Plugins string `gorm:"not null;column:plugins;type:varchar(32);size:96" json:"plugins,omitempty"`
    Sort    uint   `gorm:"not null;column:sort;type:int unsigned" json:"sort"`
    Status  int8   `gorm:"not null;column:status;type:tinyint unsigned" json:"status"`
}

func (SystemHookPlugins) TableName() string {
	return "system_hook_plugins"
}
func (this SystemHookPlugins) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemHookPlugins) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemHookPlugins) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this SystemHookPlugins) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(SystemHookPlugins{}).Error; err != nil {
		return err
	}
	return nil
}

func NewSystemHookPlugins() SystemHookPlugins {
    return SystemHookPlugins{}
}

