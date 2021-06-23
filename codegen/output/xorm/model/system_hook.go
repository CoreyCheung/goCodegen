package model

// Created by CodeMachine on 2020-05-09 22:05:35.026972 +0800 CST m=+0.138984406
// [系统] 钩子表
type SystemHook struct {
    // 创建时间
    Ctime  uint   `xorm:"'ctime' not null INT UNSIGNED" json:"ctime"`
    ID     uint   `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    // 钩子简介
    Intro  string `xorm:"'intro' not null VARCHAR(200)" json:"intro,omitempty"`
    // 更新时间
    Mtime  uint   `xorm:"'mtime' not null INT UNSIGNED" json:"mtime"`
    // 钩子名称
    Name   string `xorm:"'name' not null VARCHAR(50)" json:"name,omitempty"`
    // 钩子来源[plugins.插件名，module.模块名]
    Source string `xorm:"'source' not null VARCHAR(50)" json:"source,omitempty"`
    Status int8   `xorm:"'status' not null TINYINT UNSIGNED" json:"status"`
    // 系统插件
    System int8   `xorm:"'system' not null TINYINT UNSIGNED" json:"system"`
}

func (SystemHook) TableName() string {
	return "system_hook"
}

func NewSystemHook() SystemHook {
    return SystemHook{}
}
