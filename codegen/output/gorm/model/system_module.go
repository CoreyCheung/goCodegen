package model

// Created by CodeMachine on 2020-05-10 15:46:04.73189 +0800 CST m=+0.135204994
// [系统] 模块
type SystemModule struct {
    // 应用市场ID(0本地)
    AppId      string `gorm:"not null;column:app_id;type:varchar(30);size:90" json:"appId,omitempty"`
    // 应用秘钥
    AppKeys    string `gorm:"column:app_keys;type:varchar(200);size:600" json:"appKeys,omitempty"`
    // 作者
    Author     string `gorm:"not null;column:author;type:varchar(100);size:300" json:"author,omitempty"`
    // 配置
    Config     string `gorm:"not null;column:config;type:text;size:65535" json:"config,omitempty"`
    // 创建时间
    Ctime      uint   `gorm:"not null;column:ctime;type:int unsigned" json:"ctime"`
    // 默认模块(只能有一个)
    Default    int8   `gorm:"not null;column:default;type:tinyint unsigned" json:"default"`
    // 图标
    Icon       string `gorm:"not null;column:icon;type:varchar(80);size:240" json:"icon,omitempty"`
    ID         uint   `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    // 模块标识(模块名(字母).开发者标识.module)
    Identifier string `gorm:"not null;column:identifier;type:varchar(100);size:300" json:"identifier,omitempty"`
    // 模块简介
    Intro      string `gorm:"not null;column:intro;type:varchar(255);size:765" json:"intro,omitempty"`
    // 修改时间
    Mtime      uint   `gorm:"not null;column:mtime;type:int unsigned" json:"mtime"`
    // 模块名(英文)
    Name       string `gorm:"not null;column:name;type:varchar(50);size:150" json:"name,omitempty"`
    // 排序
    Sort       uint   `gorm:"not null;column:sort;type:int unsigned" json:"sort"`
    // 0未安装，1未启用，2已启用
    Status     int8   `gorm:"not null;column:status;type:tinyint unsigned" json:"status"`
    // 系统模块
    System     int8   `gorm:"not null;column:system;type:tinyint unsigned" json:"system"`
    // 主题模板
    Theme      string `gorm:"not null;column:theme;type:varchar(50);size:150" json:"theme,omitempty"`
    // 模块标题
    Title      string `gorm:"not null;column:title;type:varchar(50);size:150" json:"title,omitempty"`
    // 链接
    Url        string `gorm:"not null;column:url;type:varchar(255);size:765" json:"url,omitempty"`
    // 版本号
    Version    string `gorm:"not null;column:version;type:varchar(20);size:60" json:"version,omitempty"`
}

func (SystemModule) TableName() string {
	return "system_module"
}
func (this SystemModule) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemModule) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemModule) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this SystemModule) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(SystemModule{}).Error; err != nil {
		return err
	}
	return nil
}

func NewSystemModule() SystemModule {
    return SystemModule{}
}

