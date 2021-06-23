package model

// Created by CodeMachine on 2020-05-10 15:46:04.71744 +0800 CST m=+0.120755729
// [系统] 系统配置
type SystemConfig struct {
    Ctime   uint   `gorm:"not null;column:ctime;type:int unsigned" json:"ctime"`
    // 分组
    Group   string `gorm:"not null;column:group;type:varchar(20);size:60" json:"group,omitempty"`
    ID      uint   `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    Mtime   uint   `gorm:"not null;column:mtime;type:int unsigned" json:"mtime"`
    // 配置名称，由英文字母和下划线组成
    Name    string `gorm:"not null;column:name;type:varchar(50);size:150" json:"name,omitempty"`
    // 配置项(选项名:选项值)
    Options string `gorm:"not null;column:options;type:text;size:65535" json:"options,omitempty"`
    // 排序
    Sort    uint   `gorm:"not null;column:sort;type:int unsigned" json:"sort"`
    // 状态
    Status  int8   `gorm:"not null;column:status;type:tinyint unsigned" json:"status"`
    // 是否为系统配置(1是，0否)
    System  int8   `gorm:"not null;column:system;type:tinyint unsigned" json:"system"`
    // 配置提示
    Tips    string `gorm:"not null;column:tips;type:varchar(255);size:765" json:"tips,omitempty"`
    // 配置标题
    Title   string `gorm:"not null;column:title;type:varchar(20);size:60" json:"title,omitempty"`
    // 配置类型()
    Type    string `gorm:"not null;column:type;type:varchar(20);size:60" json:"type,omitempty"`
    // 文件上传接口
    Url     string `gorm:"not null;column:url;type:varchar(255);size:765" json:"url,omitempty"`
    // 配置值
    Value   string `gorm:"not null;column:value;type:text;size:65535" json:"value,omitempty"`
}

func (SystemConfig) TableName() string {
	return "system_config"
}
func (this SystemConfig) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemConfig) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemConfig) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this SystemConfig) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(SystemConfig{}).Error; err != nil {
		return err
	}
	return nil
}

func NewSystemConfig() SystemConfig {
    return SystemConfig{}
}

