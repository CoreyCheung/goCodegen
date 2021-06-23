package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.952316 +0800 CST m=+0.064330958
type CouponDetail struct {
    CardId       string         `xorm:"'card_id' VARCHAR(255)" json:"cardId,omitempty"`
    CouponTypeId uint           `xorm:"'coupon_type_id' INT UNSIGNED" json:"couponTypeId"`
    CreatedAt    date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt    *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    EndTime      date.Datetime  `xorm:"'end_time' TIMESTAMP" json:"endTime,omitempty"`
    ID           uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    Mobile       string         `xorm:"'mobile' not null VARCHAR(255)" json:"mobile,omitempty"`
    ReceiveTime  date.Datetime  `xorm:"'receive_time' TIMESTAMP" json:"receiveTime,omitempty"`
    SellPoint    int            `xorm:"'sell_point' INT" json:"sellPoint"`
    StartTime    date.Datetime  `xorm:"'start_time' TIMESTAMP" json:"startTime,omitempty"`
    Status       int            `xorm:"'status' INT" json:"status"`
    UpdatedAt    date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
    UseTime      date.Datetime  `xorm:"'use_time' TIMESTAMP" json:"useTime,omitempty"`
}

func (CouponDetail) TableName() string {
	return "coupon_detail"
}

func NewCouponDetail() CouponDetail {
    return CouponDetail{}
}
