package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.974057 +0800 CST m=+0.086070774
type GiftcardTransactionDetail struct {
    Amount       float64        `xorm:"'amount' DOUBLE" json:"amount"`
    CreatedAt    date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt    *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    Description  string         `xorm:"'description' VARCHAR(255)" json:"description,omitempty"`
    GiftCardId   int            `xorm:"'gift_card_id' INT" json:"giftCardId"`
    GiftCardName string         `xorm:"'gift_card_name' VARCHAR(255)" json:"giftCardName,omitempty"`
    ID           uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    Mobile       string         `xorm:"'mobile' VARCHAR(255)" json:"mobile,omitempty"`
    // 充值和退款的操作员
    Operator     string         `xorm:"'operator' VARCHAR(255)" json:"operator,omitempty"`
    OrderId      string         `xorm:"'order_id' VARCHAR(255)" json:"orderId,omitempty"`
    ShopId       int            `xorm:"'shop_id' INT" json:"shopId"`
    Type         int            `xorm:"'type' INT" json:"type"`
    UpdatedAt    date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (GiftcardTransactionDetail) TableName() string {
	return "giftcard_transaction_detail"
}

func NewGiftcardTransactionDetail() GiftcardTransactionDetail {
    return GiftcardTransactionDetail{}
}
