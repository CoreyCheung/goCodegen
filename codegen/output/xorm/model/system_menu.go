package model

// Created by CodeMachine on 2020-05-09 22:05:35.037374 +0800 CST m=+0.149386561
// [系统] 管理菜单
type SystemMenu struct {
    Ctime  uint   `xorm:"'ctime' not null INT UNSIGNED" json:"ctime"`
    // 开发模式可见
    Debug  int8   `xorm:"'debug' not null TINYINT UNSIGNED" json:"debug"`
    // 菜单图标
    Icon   string `xorm:"'icon' not null VARCHAR(80)" json:"icon,omitempty"`
    ID     uint   `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    // 模块名或插件名，插件名格式:plugins.插件名
    Module string `xorm:"'module' not null VARCHAR(20)" json:"module,omitempty"`
    // 是否为菜单显示，1显示0不显示
    Nav    int8   `xorm:"'nav' not null TINYINT UNSIGNED" json:"nav"`
    // 扩展参数
    Param  string `xorm:"'param' not null VARCHAR(200)" json:"param,omitempty"`
    Pid    uint   `xorm:"'pid' not null INT UNSIGNED" json:"pid"`
    // 排序
    Sort   uint   `xorm:"'sort' not null INT UNSIGNED" json:"sort"`
    // 状态1显示，0隐藏
    Status int8   `xorm:"'status' not null TINYINT UNSIGNED" json:"status"`
    // 是否为系统菜单，系统菜单不可删除
    System int8   `xorm:"'system' not null TINYINT UNSIGNED" json:"system"`
    // 打开方式(_blank,_self)
    Target string `xorm:"'target' not null VARCHAR(20)" json:"target,omitempty"`
    // 菜单标题
    Title  string `xorm:"'title' not null VARCHAR(20)" json:"title,omitempty"`
    // 管理员ID(快捷菜单专用)
    Uid    uint   `xorm:"'uid' not null INT UNSIGNED" json:"uid"`
    // 链接地址(模块/控制器/方法)
    Url    string `xorm:"'url' not null VARCHAR(200)" json:"url,omitempty"`
}

func (SystemMenu) TableName() string {
	return "system_menu"
}

func NewSystemMenu() SystemMenu {
    return SystemMenu{}
}
