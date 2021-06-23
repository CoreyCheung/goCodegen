package model

// Created by CodeMachine on 2020-05-09 22:05:35.019231 +0800 CST m=+0.131243373
// [系统] 系统配置
type SystemConfig struct {
    Ctime   uint   `xorm:"'ctime' not null INT UNSIGNED" json:"ctime"`
    // 分组
    Group   string `xorm:"'group' not null VARCHAR(20)" json:"group,omitempty"`
    ID      uint   `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    Mtime   uint   `xorm:"'mtime' not null INT UNSIGNED" json:"mtime"`
    // 配置名称，由英文字母和下划线组成
    Name    string `xorm:"'name' not null VARCHAR(50)" json:"name,omitempty"`
    // 配置项(选项名:选项值)
    Options string `xorm:"'options' not null TEXT" json:"options,omitempty"`
    // 排序
    Sort    uint   `xorm:"'sort' not null INT UNSIGNED" json:"sort"`
    // 状态
    Status  int8   `xorm:"'status' not null TINYINT UNSIGNED" json:"status"`
    // 是否为系统配置(1是，0否)
    System  int8   `xorm:"'system' not null TINYINT UNSIGNED" json:"system"`
    // 配置提示
    Tips    string `xorm:"'tips' not null VARCHAR(255)" json:"tips,omitempty"`
    // 配置标题
    Title   string `xorm:"'title' not null VARCHAR(20)" json:"title,omitempty"`
    // 配置类型()
    Type    string `xorm:"'type' not null VARCHAR(20)" json:"type,omitempty"`
    // 文件上传接口
    Url     string `xorm:"'url' not null VARCHAR(255)" json:"url,omitempty"`
    // 配置值
    Value   string `xorm:"'value' not null TEXT" json:"value,omitempty"`
}

func (SystemConfig) TableName() string {
	return "system_config"
}

func NewSystemConfig() SystemConfig {
    return SystemConfig{}
}
