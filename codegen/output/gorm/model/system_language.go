package model

// Created by CodeMachine on 2020-05-10 15:46:04.723235 +0800 CST m=+0.126550845
// [系统] 语言包
type SystemLanguage struct {
    // 编码
    Code   string `gorm:"not null;column:code;type:varchar(20);size:60" json:"code,omitempty"`
    // 图标
    Icon   string `gorm:"not null;column:icon;type:varchar(30);size:90" json:"icon,omitempty"`
    ID     uint   `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    // 本地浏览器语言编码
    Locale string `gorm:"not null;column:locale;type:varchar(255);size:765" json:"locale,omitempty"`
    // 语言包名称
    Name   string `gorm:"not null;column:name;type:varchar(50);size:150" json:"name,omitempty"`
    // 上传的语言包
    Pack   string `gorm:"not null;column:pack;type:varchar(100);size:300" json:"pack,omitempty"`
    Sort   int8   `gorm:"not null;column:sort;type:tinyint unsigned" json:"sort"`
    Status int8   `gorm:"not null;column:status;type:tinyint unsigned" json:"status"`
}

func (SystemLanguage) TableName() string {
	return "system_language"
}
func (this SystemLanguage) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemLanguage) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemLanguage) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this SystemLanguage) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(SystemLanguage{}).Error; err != nil {
		return err
	}
	return nil
}

func NewSystemLanguage() SystemLanguage {
    return SystemLanguage{}
}

