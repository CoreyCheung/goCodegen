package model

// Created by CodeMachine on 2020-05-10 15:46:04.734245 +0800 CST m=+0.137560209
// [系统] 插件表
type SystemPlugins struct {
    // 来源(0本地)
    AppId      string `gorm:"not null;column:app_id;type:varchar(30);size:90" json:"appId,omitempty"`
    // 应用秘钥
    AppKeys    string `gorm:"column:app_keys;type:varchar(200);size:600" json:"appKeys,omitempty"`
    // 作者
    Author     string `gorm:"not null;column:author;type:varchar(32);size:96" json:"author,omitempty"`
    // 插件配置
    Config     string `gorm:"not null;column:config;type:text;size:65535" json:"config,omitempty"`
    Ctime      uint   `gorm:"not null;column:ctime;type:int unsigned" json:"ctime"`
    // 图标
    Icon       string `gorm:"not null;column:icon;type:varchar(64);size:192" json:"icon,omitempty"`
    ID         uint   `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    // 插件唯一标识符
    Identifier string `gorm:"not null;column:identifier;type:varchar(64);size:192" json:"identifier,omitempty"`
    // 插件简介
    Intro      string `gorm:"not null;column:intro;type:text;size:65535" json:"intro,omitempty"`
    Mtime      uint   `gorm:"not null;column:mtime;type:int unsigned" json:"mtime"`
    // 插件名称(英文)
    Name       string `gorm:"not null;column:name;type:varchar(32);size:96" json:"name,omitempty"`
    // 排序
    Sort       uint   `gorm:"not null;column:sort;type:int unsigned" json:"sort"`
    // 状态
    Status     int8   `gorm:"not null;column:status;type:tinyint unsigned" json:"status"`
    System     int8   `gorm:"not null;column:system;type:tinyint unsigned" json:"system"`
    // 插件标题
    Title      string `gorm:"not null;column:title;type:varchar(32);size:96" json:"title,omitempty"`
    // 作者主页
    Url        string `gorm:"not null;column:url;type:varchar(255);size:765" json:"url,omitempty"`
    // 版本号
    Version    string `gorm:"not null;column:version;type:varchar(16);size:48" json:"version,omitempty"`
}

func (SystemPlugins) TableName() string {
	return "system_plugins"
}
func (this SystemPlugins) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemPlugins) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemPlugins) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this SystemPlugins) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(SystemPlugins{}).Error; err != nil {
		return err
	}
	return nil
}

func NewSystemPlugins() SystemPlugins {
    return SystemPlugins{}
}

