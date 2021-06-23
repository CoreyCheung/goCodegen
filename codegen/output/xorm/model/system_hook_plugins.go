package model

// Created by CodeMachine on 2020-05-09 22:05:35.029135 +0800 CST m=+0.141147323
// [系统] 钩子-插件对应表
type SystemHookPlugins struct {
    Ctime   uint   `xorm:"'ctime' not null INT UNSIGNED" json:"ctime"`
    // 钩子id
    Hook    string `xorm:"'hook' not null VARCHAR(32)" json:"hook,omitempty"`
    ID      uint   `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    Mtime   uint   `xorm:"'mtime' not null INT UNSIGNED" json:"mtime"`
    // 插件标识
    Plugins string `xorm:"'plugins' not null VARCHAR(32)" json:"plugins,omitempty"`
    Sort    uint   `xorm:"'sort' not null INT UNSIGNED" json:"sort"`
    Status  int8   `xorm:"'status' not null TINYINT UNSIGNED" json:"status"`
}

func (SystemHookPlugins) TableName() string {
	return "system_hook_plugins"
}

func NewSystemHookPlugins() SystemHookPlugins {
    return SystemHookPlugins{}
}
