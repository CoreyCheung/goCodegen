package model

// Created by CodeMachine on 2020-05-09 22:05:35.013548 +0800 CST m=+0.125561262
// [系统] 上传附件
type SystemAnnex struct {
    Ctime  uint    `xorm:"'ctime' not null INT UNSIGNED" json:"ctime"`
    // 关联的数据ID
    DataId uint    `xorm:"'data_id' not null INT UNSIGNED" json:"dataId"`
    // 上传文件
    File   string  `xorm:"'file' not null VARCHAR(255)" json:"file,omitempty"`
    // 文件分组
    Group  string  `xorm:"'group' not null VARCHAR(100)" json:"group,omitempty"`
    // 文件hash值
    Hash   string  `xorm:"'hash' not null VARCHAR(64)" json:"hash,omitempty"`
    ID     uint    `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    // 附件大小KB
    Size   float64 `xorm:"'size' not null DECIMAL(12,2) UNSIGNED" json:"size"`
    // 使用状态(0未使用，1已使用)
    Status int8    `xorm:"'status' not null TINYINT UNSIGNED" json:"status"`
    // 类型
    Type   string  `xorm:"'type' not null VARCHAR(20)" json:"type,omitempty"`
}

func (SystemAnnex) TableName() string {
	return "system_annex"
}

func NewSystemAnnex() SystemAnnex {
    return SystemAnnex{}
}
