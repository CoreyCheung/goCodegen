package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.653512 +0800 CST m=+0.056829076
type CouponRanges struct {
    CouponTypeId int            `gorm:"column:coupon_type_id;type:int" json:"couponTypeId"`
    CreatedAt    date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt    *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    ID           uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    ShopId       int            `gorm:"column:shop_id;type:int" json:"shopId"`
    ShopName     string         `gorm:"column:shop_name;type:varchar(255);size:1020" json:"shopName,omitempty"`
    UpdatedAt    date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (CouponRanges) TableName() string {
	return "coupon_ranges"
}
func (this CouponRanges) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this CouponRanges) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this CouponRanges) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this CouponRanges) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(CouponRanges{}).Error; err != nil {
		return err
	}
	return nil
}

func NewCouponRanges() CouponRanges {
    return CouponRanges{}
}

