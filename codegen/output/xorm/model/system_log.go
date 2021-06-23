package model

// Created by CodeMachine on 2020-05-09 22:05:35.035665 +0800 CST m=+0.147677712
// [系统] 操作日志
type SystemLog struct {
    Count  uint   `xorm:"'count' not null INT UNSIGNED" json:"count"`
    Ctime  uint   `xorm:"'ctime' not null INT UNSIGNED" json:"ctime"`
    ID     uint   `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    Ip     string `xorm:"'ip' VARCHAR(128)" json:"ip,omitempty"`
    Mtime  uint   `xorm:"'mtime' not null INT UNSIGNED" json:"mtime"`
    Param  string `xorm:"'param' TEXT" json:"param,omitempty"`
    Remark string `xorm:"'remark' VARCHAR(255)" json:"remark,omitempty"`
    Title  string `xorm:"'title' VARCHAR(100)" json:"title,omitempty"`
    Uid    uint   `xorm:"'uid' not null INT UNSIGNED" json:"uid"`
    Url    string `xorm:"'url' VARCHAR(200)" json:"url,omitempty"`
}

func (SystemLog) TableName() string {
	return "system_log"
}

func NewSystemLog() SystemLog {
    return SystemLog{}
}
