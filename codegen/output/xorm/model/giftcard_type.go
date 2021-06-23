package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.97518 +0800 CST m=+0.087193906
type GiftcardType struct {
    CardId         string         `xorm:"'card_id' VARCHAR(255)" json:"cardId,omitempty"`
    CompanyLogo    string         `xorm:"'company_logo' VARCHAR(255)" json:"companyLogo,omitempty"`
    CompanyName    string         `xorm:"'company_name' VARCHAR(255)" json:"companyName,omitempty"`
    CreatedAt      date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt      *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    Description    string         `xorm:"'description' VARCHAR(255)" json:"description,omitempty"`
    // 0不限制
    ExpiryDateType int            `xorm:"'expiry_date_type' INT" json:"expiryDateType"`
    GiftCardName   string         `xorm:"'gift_card_name' VARCHAR(255)" json:"giftCardName,omitempty"`
    ID             uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    IssuedCount    int            `xorm:"'issued_count' INT" json:"issuedCount"`
    RegularDate    int            `xorm:"'regular_date' INT" json:"regularDate"`
    TotalAmount    float64        `xorm:"'total_amount' DOUBLE" json:"totalAmount"`
    UpdatedAt      date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
    UseEndTime     date.Datetime  `xorm:"'use_end_time' TIMESTAMP" json:"useEndTime,omitempty"`
    UseRangeType   int            `xorm:"'use_range_type' INT" json:"useRangeType"`
}

func (GiftcardType) TableName() string {
	return "giftcard_type"
}

func NewGiftcardType() GiftcardType {
    return GiftcardType{}
}
