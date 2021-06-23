package model

// Created by CodeMachine on 2020-05-09 22:05:35.040004 +0800 CST m=+0.152016181
// [系统] 管理菜单语言包
type SystemMenuLang struct {
    ID     uint   `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    // 语言包
    Lang   int8   `xorm:"'lang' not null TINYINT UNSIGNED" json:"lang"`
    MenuId uint   `xorm:"'menu_id' not null INT UNSIGNED" json:"menuId"`
    // 标题
    Title  string `xorm:"'title' not null VARCHAR(120)" json:"title,omitempty"`
}

func (SystemMenuLang) TableName() string {
	return "system_menu_lang"
}

func NewSystemMenuLang() SystemMenuLang {
    return SystemMenuLang{}
}
