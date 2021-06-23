package model

// Created by CodeMachine on 2020-05-09 22:05:35.045405 +0800 CST m=+0.157416978
// [系统] 插件表
type SystemPlugins struct {
    // 来源(0本地)
    AppId      string `xorm:"'app_id' not null VARCHAR(30)" json:"appId,omitempty"`
    // 应用秘钥
    AppKeys    string `xorm:"'app_keys' VARCHAR(200)" json:"appKeys,omitempty"`
    // 作者
    Author     string `xorm:"'author' not null VARCHAR(32)" json:"author,omitempty"`
    // 插件配置
    Config     string `xorm:"'config' not null TEXT" json:"config,omitempty"`
    Ctime      uint   `xorm:"'ctime' not null INT UNSIGNED" json:"ctime"`
    // 图标
    Icon       string `xorm:"'icon' not null VARCHAR(64)" json:"icon,omitempty"`
    ID         uint   `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    // 插件唯一标识符
    Identifier string `xorm:"'identifier' not null VARCHAR(64)" json:"identifier,omitempty"`
    // 插件简介
    Intro      string `xorm:"'intro' not null TEXT" json:"intro,omitempty"`
    Mtime      uint   `xorm:"'mtime' not null INT UNSIGNED" json:"mtime"`
    // 插件名称(英文)
    Name       string `xorm:"'name' not null VARCHAR(32)" json:"name,omitempty"`
    // 排序
    Sort       uint   `xorm:"'sort' not null INT UNSIGNED" json:"sort"`
    // 状态
    Status     int8   `xorm:"'status' not null TINYINT UNSIGNED" json:"status"`
    System     int8   `xorm:"'system' not null TINYINT UNSIGNED" json:"system"`
    // 插件标题
    Title      string `xorm:"'title' not null VARCHAR(32)" json:"title,omitempty"`
    // 作者主页
    Url        string `xorm:"'url' not null VARCHAR(255)" json:"url,omitempty"`
    // 版本号
    Version    string `xorm:"'version' not null VARCHAR(16)" json:"version,omitempty"`
}

func (SystemPlugins) TableName() string {
	return "system_plugins"
}

func NewSystemPlugins() SystemPlugins {
    return SystemPlugins{}
}
