package model

// Created by CodeMachine on 2020-05-10 15:46:04.727379 +0800 CST m=+0.130694496
// [系统] 管理菜单
type SystemMenu struct {
    Ctime  uint   `gorm:"not null;column:ctime;type:int unsigned" json:"ctime"`
    // 开发模式可见
    Debug  int8   `gorm:"not null;column:debug;type:tinyint unsigned" json:"debug"`
    // 菜单图标
    Icon   string `gorm:"not null;column:icon;type:varchar(80);size:240" json:"icon,omitempty"`
    ID     uint   `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    // 模块名或插件名，插件名格式:plugins.插件名
    Module string `gorm:"not null;column:module;type:varchar(20);size:60" json:"module,omitempty"`
    // 是否为菜单显示，1显示0不显示
    Nav    int8   `gorm:"not null;column:nav;type:tinyint unsigned" json:"nav"`
    // 扩展参数
    Param  string `gorm:"not null;column:param;type:varchar(200);size:600" json:"param,omitempty"`
    Pid    uint   `gorm:"not null;column:pid;type:int unsigned" json:"pid"`
    // 排序
    Sort   uint   `gorm:"not null;column:sort;type:int unsigned" json:"sort"`
    // 状态1显示，0隐藏
    Status int8   `gorm:"not null;column:status;type:tinyint unsigned" json:"status"`
    // 是否为系统菜单，系统菜单不可删除
    System int8   `gorm:"not null;column:system;type:tinyint unsigned" json:"system"`
    // 打开方式(_blank,_self)
    Target string `gorm:"not null;column:target;type:varchar(20);size:60" json:"target,omitempty"`
    // 菜单标题
    Title  string `gorm:"not null;column:title;type:varchar(20);size:60" json:"title,omitempty"`
    // 管理员ID(快捷菜单专用)
    Uid    uint   `gorm:"not null;column:uid;type:int unsigned" json:"uid"`
    // 链接地址(模块/控制器/方法)
    Url    string `gorm:"not null;column:url;type:varchar(200);size:600" json:"url,omitempty"`
}

func (SystemMenu) TableName() string {
	return "system_menu"
}
func (this SystemMenu) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemMenu) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemMenu) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this SystemMenu) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(SystemMenu{}).Error; err != nil {
		return err
	}
	return nil
}

func NewSystemMenu() SystemMenu {
    return SystemMenu{}
}

