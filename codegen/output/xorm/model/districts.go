package model

// Created by CodeMachine on 2020-05-09 22:05:34.96273 +0800 CST m=+0.074744463
type Districts struct {
    // 0:不展示 1：展示
    Bshow        int    `xorm:"'bshow' INT" json:"bshow"`
    // 城市的名字
    DistrictName string `xorm:"'district_name' VARCHAR(255)" json:"districtName,omitempty"`
    // 层级序列
    DistrictSqe  string `xorm:"'district_sqe' VARCHAR(255)" json:"districtSqe,omitempty"`
    // 地区所处的层级
    Hierarchy    int    `xorm:"'hierarchy' INT" json:"hierarchy"`
    // 1:热门地区/城市
    Hot          int    `xorm:"'hot' INT" json:"hot"`
    // 主键自增
    ID           uint   `xorm:"'id' not null pk INT UNSIGNED" json:"id"`
    // 父类id
    Pid          int    `xorm:"'pid' INT" json:"pid"`
    // 城市的类型，0是国，1是省，2是市，3是区
    Type         int    `xorm:"'type' INT" json:"type"`
}

func (Districts) TableName() string {
	return "districts"
}

func NewDistricts() Districts {
    return Districts{}
}
