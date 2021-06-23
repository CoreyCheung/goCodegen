package model

// Created by CodeMachine on 2020-05-09 22:05:35.018125 +0800 CST m=+0.130137764
// [系统] 附件分组
type SystemAnnexGroup struct {
    // 附件数量
    Count uint    `xorm:"'count' not null INT UNSIGNED" json:"count"`
    ID    uint    `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    // 附件分组
    Name  string  `xorm:"'name' not null VARCHAR(50)" json:"name,omitempty"`
    // 附件大小kb
    Size  float64 `xorm:"'size' not null DECIMAL(12,2)" json:"size"`
}

func (SystemAnnexGroup) TableName() string {
	return "system_annex_group"
}

func NewSystemAnnexGroup() SystemAnnexGroup {
    return SystemAnnexGroup{}
}
