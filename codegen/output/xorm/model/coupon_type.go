package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.956392 +0800 CST m=+0.068406556
type CouponType struct {
    ActivitiesId      int           `xorm:"'activities_id' INT" json:"activitiesId"`
    BenefitType       int           `xorm:"'benefit_type' INT" json:"benefitType"`
    CouponName        string        `xorm:"'coupon_name' VARCHAR(255)" json:"couponName,omitempty"`
    Decrease          float64       `xorm:"'decrease' DOUBLE" json:"decrease"`
    Description       string        `xorm:"'description' VARCHAR(255)" json:"description,omitempty"`
    DiscountRate      float64       `xorm:"'discount_rate' DOUBLE" json:"discountRate"`
    EachReceiveCount  int           `xorm:"'each_receive_count' INT" json:"eachReceiveCount"`
    ExpiryDateType    int           `xorm:"'expiry_date_type' INT" json:"expiryDateType"`
    ID                uint          `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    MinCharge         float64       `xorm:"'min_charge' DOUBLE" json:"minCharge"`
    Points            int           `xorm:"'points' INT" json:"points"`
    ReceivableBalance int           `xorm:"'receivable_balance' INT" json:"receivableBalance"`
    ReceivableTotal   int           `xorm:"'receivable_total' INT" json:"receivableTotal"`
    ReceiveEndTime    date.Datetime `xorm:"'receive_end_time' TIMESTAMP" json:"receiveEndTime,omitempty"`
    ReceiveStartTime  date.Datetime `xorm:"'receive_start_time' TIMESTAMP" json:"receiveStartTime,omitempty"`
    ReceiveStatus     int           `xorm:"'receive_status' INT" json:"receiveStatus"`
    ReceiveType       int           `xorm:"'receive_type' INT" json:"receiveType"`
    RegularDate       int           `xorm:"'regular_date' INT" json:"regularDate"`
    Transferable      int           `xorm:"'transferable' INT" json:"transferable"`
    UseEndTime        date.Datetime `xorm:"'use_end_time' TIMESTAMP" json:"useEndTime,omitempty"`
    UseRangeType      int           `xorm:"'use_range_type' INT" json:"useRangeType"`
    UseStartTime      date.Datetime `xorm:"'use_start_time' TIMESTAMP" json:"useStartTime,omitempty"`
}

func (CouponType) TableName() string {
	return "coupon_type"
}

func NewCouponType() CouponType {
    return CouponType{}
}
