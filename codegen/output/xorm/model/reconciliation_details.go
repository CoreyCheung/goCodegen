package model

import "git.wanpinghui.com/WPH/go_common/wph"
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.997327 +0800 CST m=+0.109340886
type ReconciliationDetails struct {
    CreatedAt          date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt          *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    ID                 uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    InconsistentReason string         `xorm:"'inconsistent_reason' VARCHAR(255)" json:"inconsistentReason,omitempty"`
    OrderId            string         `xorm:"'order_id' VARCHAR(255)" json:"orderId,omitempty"`
    PaymentType        string         `xorm:"'payment_type' VARCHAR(255)" json:"paymentType,omitempty"`
    States             wph.Long       `xorm:"'states' BIGINT" json:"states"`
    UpdatedAt          date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (ReconciliationDetails) TableName() string {
	return "reconciliation_details"
}

func NewReconciliationDetails() ReconciliationDetails {
    return ReconciliationDetails{}
}
