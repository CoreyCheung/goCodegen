package model

// Created by CodeMachine on 2020-05-09 22:05:35.041064 +0800 CST m=+0.153076454
// [系统] 模块
type SystemModule struct {
    // 应用市场ID(0本地)
    AppId      string `xorm:"'app_id' not null VARCHAR(30)" json:"appId,omitempty"`
    // 应用秘钥
    AppKeys    string `xorm:"'app_keys' VARCHAR(200)" json:"appKeys,omitempty"`
    // 作者
    Author     string `xorm:"'author' not null VARCHAR(100)" json:"author,omitempty"`
    // 配置
    Config     string `xorm:"'config' not null TEXT" json:"config,omitempty"`
    // 创建时间
    Ctime      uint   `xorm:"'ctime' not null INT UNSIGNED" json:"ctime"`
    // 默认模块(只能有一个)
    Default    int8   `xorm:"'default' not null TINYINT UNSIGNED" json:"default"`
    // 图标
    Icon       string `xorm:"'icon' not null VARCHAR(80)" json:"icon,omitempty"`
    ID         uint   `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    // 模块标识(模块名(字母).开发者标识.module)
    Identifier string `xorm:"'identifier' not null VARCHAR(100)" json:"identifier,omitempty"`
    // 模块简介
    Intro      string `xorm:"'intro' not null VARCHAR(255)" json:"intro,omitempty"`
    // 修改时间
    Mtime      uint   `xorm:"'mtime' not null INT UNSIGNED" json:"mtime"`
    // 模块名(英文)
    Name       string `xorm:"'name' not null VARCHAR(50)" json:"name,omitempty"`
    // 排序
    Sort       uint   `xorm:"'sort' not null INT UNSIGNED" json:"sort"`
    // 0未安装，1未启用，2已启用
    Status     int8   `xorm:"'status' not null TINYINT UNSIGNED" json:"status"`
    // 系统模块
    System     int8   `xorm:"'system' not null TINYINT UNSIGNED" json:"system"`
    // 主题模板
    Theme      string `xorm:"'theme' not null VARCHAR(50)" json:"theme,omitempty"`
    // 模块标题
    Title      string `xorm:"'title' not null VARCHAR(50)" json:"title,omitempty"`
    // 链接
    Url        string `xorm:"'url' not null VARCHAR(255)" json:"url,omitempty"`
    // 版本号
    Version    string `xorm:"'version' not null VARCHAR(20)" json:"version,omitempty"`
}

func (SystemModule) TableName() string {
	return "system_module"
}

func NewSystemModule() SystemModule {
    return SystemModule{}
}
