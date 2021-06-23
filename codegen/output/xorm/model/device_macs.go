package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.96203 +0800 CST m=+0.074044291
type DeviceMacs struct {
    CreatedAt date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    DepId     int            `xorm:"'depId' INT" json:"depId"`
    DeviceMac string         `xorm:"'deviceMac' VARCHAR(255)" json:"deviceMac,omitempty"`
    ID        uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    ShopId    uint           `xorm:"'shop_id' INT UNSIGNED" json:"shopId"`
    UpdatedAt date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (DeviceMacs) TableName() string {
	return "device_macs"
}

func NewDeviceMacs() DeviceMacs {
    return DeviceMacs{}
}
