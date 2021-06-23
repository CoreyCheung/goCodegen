package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.95456 +0800 CST m=+0.066574876
type CouponPointTransactionDetail struct {
    CouponId         int            `xorm:"'coupon_id' INT" json:"couponId"`
    CouponName       string         `xorm:"'coupon_name' VARCHAR(255)" json:"couponName,omitempty"`
    CreatedAt        date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt        *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    Description      string         `xorm:"'description' VARCHAR(255)" json:"description,omitempty"`
    ID               uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    Mobile           string         `xorm:"'mobile' VARCHAR(255)" json:"mobile,omitempty"`
    OrderId          string         `xorm:"'order_id' VARCHAR(255)" json:"orderId,omitempty"`
    Points           int            `xorm:"'points' INT" json:"points"`
    ShopId           int            `xorm:"'shop_id' INT" json:"shopId"`
    TargetUserId     string         `xorm:"'target_user_id' VARCHAR(255)" json:"targetUserId,omitempty"`
    TargetUserMobile string         `xorm:"'target_user_mobile' VARCHAR(255)" json:"targetUserMobile,omitempty"`
    Type             int            `xorm:"'type' INT" json:"type"`
    UpdatedAt        date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (CouponPointTransactionDetail) TableName() string {
	return "coupon_point_transaction_detail"
}

func NewCouponPointTransactionDetail() CouponPointTransactionDetail {
    return CouponPointTransactionDetail{}
}
