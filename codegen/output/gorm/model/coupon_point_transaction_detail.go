package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.65147 +0800 CST m=+0.054787675
type CouponPointTransactionDetail struct {
    CouponId         int            `gorm:"column:coupon_id;type:int" json:"couponId"`
    CouponName       string         `gorm:"column:coupon_name;type:varchar(255);size:1020" json:"couponName,omitempty"`
    CreatedAt        date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt        *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    Description      string         `gorm:"column:description;type:varchar(255);size:1020" json:"description,omitempty"`
    ID               uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    Mobile           string         `gorm:"column:mobile;type:varchar(255);size:1020" json:"mobile,omitempty"`
    OrderId          string         `gorm:"column:order_id;type:varchar(255);size:1020" json:"orderId,omitempty"`
    Points           int            `gorm:"column:points;type:int" json:"points"`
    ShopId           int            `gorm:"column:shop_id;type:int" json:"shopId"`
    TargetUserId     string         `gorm:"column:target_user_id;type:varchar(255);size:1020" json:"targetUserId,omitempty"`
    TargetUserMobile string         `gorm:"column:target_user_mobile;type:varchar(255);size:1020" json:"targetUserMobile,omitempty"`
    Type             int            `gorm:"column:type;type:int" json:"type"`
    UpdatedAt        date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (CouponPointTransactionDetail) TableName() string {
	return "coupon_point_transaction_detail"
}
func (this CouponPointTransactionDetail) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this CouponPointTransactionDetail) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this CouponPointTransactionDetail) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this CouponPointTransactionDetail) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(CouponPointTransactionDetail{}).Error; err != nil {
		return err
	}
	return nil
}

func NewCouponPointTransactionDetail() CouponPointTransactionDetail {
    return CouponPointTransactionDetail{}
}

