package model

// Created by CodeMachine on 2020-05-09 22:05:35.033742 +0800 CST m=+0.145754643
// [系统] 语言包
type SystemLanguage struct {
    // 编码
    Code   string `xorm:"'code' not null VARCHAR(20)" json:"code,omitempty"`
    // 图标
    Icon   string `xorm:"'icon' not null VARCHAR(30)" json:"icon,omitempty"`
    ID     uint   `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    // 本地浏览器语言编码
    Locale string `xorm:"'locale' not null VARCHAR(255)" json:"locale,omitempty"`
    // 语言包名称
    Name   string `xorm:"'name' not null VARCHAR(50)" json:"name,omitempty"`
    // 上传的语言包
    Pack   string `xorm:"'pack' not null VARCHAR(100)" json:"pack,omitempty"`
    Sort   int8   `xorm:"'sort' not null TINYINT UNSIGNED" json:"sort"`
    Status int8   `xorm:"'status' not null TINYINT UNSIGNED" json:"status"`
}

func (SystemLanguage) TableName() string {
	return "system_language"
}

func NewSystemLanguage() SystemLanguage {
    return SystemLanguage{}
}
