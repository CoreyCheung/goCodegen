package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.972257 +0800 CST m=+0.084271082
type GiftcardRange struct {
    CreatedAt      date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt      *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    GiftCardTypeId uint           `xorm:"'gift_card_type_id' INT UNSIGNED" json:"giftCardTypeId"`
    ID             uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    ShopId         int            `xorm:"'shop_id' INT" json:"shopId"`
    ShopName       string         `xorm:"'shop_name' VARCHAR(255)" json:"shopName,omitempty"`
    UpdatedAt      date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (GiftcardRange) TableName() string {
	return "giftcard_range"
}

func NewGiftcardRange() GiftcardRange {
    return GiftcardRange{}
}
