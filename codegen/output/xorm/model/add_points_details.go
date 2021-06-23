package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 21:34:21.034941 +0800 CST m=+0.012805119
type AddPointsDetails struct {
    AccountAddress string         `xorm:"'account_address' VARCHAR(255)" json:"accountAddress,omitempty"`
    AddPoints      int            `xorm:"'add_points' INT" json:"addPoints"`
    CreatedAt      date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt      *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    ID             uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    OrderId        string         `xorm:"'order_id' VARCHAR(255)" json:"orderId,omitempty"`
    ProjectId      int            `xorm:"'project_id' INT" json:"projectId"`
    TransactionId  string         `xorm:"'transaction_id' VARCHAR(255)" json:"transactionId,omitempty"`
    UpdatedAt      date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (AddPointsDetails) TableName() string {
	return "add_points_details"
}

func NewAddPointsDetails() AddPointsDetails {
    return AddPointsDetails{}
}
