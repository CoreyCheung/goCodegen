package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:35.001683 +0800 CST m=+0.113696648
type SettlementDetails struct {
    Amount    float64        `xorm:"'amount' not null DOUBLE" json:"amount"`
    CreatedAt date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    ID        uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    OrderId   string         `xorm:"'order_id' not null VARCHAR(255)" json:"orderId,omitempty"`
    UpdatedAt date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (SettlementDetails) TableName() string {
	return "settlement_details"
}

func NewSettlementDetails() SettlementDetails {
    return SettlementDetails{}
}
