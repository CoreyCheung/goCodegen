package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.95349 +0800 CST m=+0.065504485
type CouponModels struct {
    BenefitType    int           `xorm:"'benefit_type' INT" json:"benefitType"`
    CouponName     string        `xorm:"'coupon_name' VARCHAR(255)" json:"couponName,omitempty"`
    Decrease       float64       `xorm:"'decrease' DOUBLE" json:"decrease"`
    Description    string        `xorm:"'description' VARCHAR(255)" json:"description,omitempty"`
    DiscountRate   float64       `xorm:"'discount_rate' DOUBLE" json:"discountRate"`
    ExpiryDateType int           `xorm:"'expiry_date_type' INT" json:"expiryDateType"`
    MinCharge      float64       `xorm:"'min_charge' DOUBLE" json:"minCharge"`
    ModelId        uint          `xorm:"'model_id' not null pk autoincr INT UNSIGNED" json:"modelId"`
    ReceiveType    int           `xorm:"'receive_type' INT" json:"receiveType"`
    RegularDate    int           `xorm:"'regular_date' INT" json:"regularDate"`
    Transferable   int           `xorm:"'transferable' INT" json:"transferable"`
    UseEndTime     date.Datetime `xorm:"'use_end_time' TIMESTAMP" json:"useEndTime,omitempty"`
    UseRangeType   int           `xorm:"'use_range_type' INT" json:"useRangeType"`
    UseStartTime   date.Datetime `xorm:"'use_start_time' TIMESTAMP" json:"useStartTime,omitempty"`
}

func (CouponModels) TableName() string {
	return "coupon_models"
}

func NewCouponModels() CouponModels {
    return CouponModels{}
}
