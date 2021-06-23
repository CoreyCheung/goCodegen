package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.969959 +0800 CST m=+0.081973367
type GiftcardDetail struct {
    CardId         string         `xorm:"'card_id' VARCHAR(255)" json:"cardId,omitempty"`
    CreatedAt      date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt      *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    EndTime        date.Datetime  `xorm:"'end_time' TIMESTAMP" json:"endTime,omitempty"`
    GiftCardTypeId uint           `xorm:"'gift_card_type_id' INT UNSIGNED" json:"giftCardTypeId"`
    ID             uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    Mobile         string         `xorm:"'mobile' not null VARCHAR(255)" json:"mobile,omitempty"`
    ReceiveTime    date.Datetime  `xorm:"'receive_time' TIMESTAMP" json:"receiveTime,omitempty"`
    RemainAmount   float64        `xorm:"'remain_amount' DOUBLE" json:"remainAmount"`
    Status         int            `xorm:"'status' not null INT" json:"status"`
    TotalAmount    float64        `xorm:"'total_amount' DOUBLE" json:"totalAmount"`
    UpdatedAt      date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (GiftcardDetail) TableName() string {
	return "giftcard_detail"
}

func NewGiftcardDetail() GiftcardDetail {
    return GiftcardDetail{}
}
