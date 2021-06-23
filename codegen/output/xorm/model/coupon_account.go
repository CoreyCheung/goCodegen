package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.949717 +0800 CST m=+0.061732122
type CouponAccount struct {
    Contributions int            `xorm:"'contributions' INT" json:"contributions"`
    CreatedAt     date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt     *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    ID            uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    Mobile        string         `xorm:"'mobile' VARCHAR(255)" json:"mobile,omitempty"`
    Points        int            `xorm:"'points' INT" json:"points"`
    UpdatedAt     date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (CouponAccount) TableName() string {
	return "coupon_account"
}

func NewCouponAccount() CouponAccount {
    return CouponAccount{}
}
