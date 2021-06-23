package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.649303 +0800 CST m=+0.052620005
type CouponModels struct {
    BenefitType    int           `gorm:"column:benefit_type;type:int" json:"benefitType"`
    CouponName     string        `gorm:"column:coupon_name;type:varchar(255);size:1020" json:"couponName,omitempty"`
    Decrease       float64       `gorm:"column:decrease;type:double" json:"decrease"`
    Description    string        `gorm:"column:description;type:varchar(255);size:1020" json:"description,omitempty"`
    DiscountRate   float64       `gorm:"column:discount_rate;type:double" json:"discountRate"`
    ExpiryDateType int           `gorm:"column:expiry_date_type;type:int" json:"expiryDateType"`
    MinCharge      float64       `gorm:"column:min_charge;type:double" json:"minCharge"`
    ModelId        uint          `gorm:"primary_key;AUTO_INCREMENT;not null;column:model_id;type:int unsigned" json:"modelId"`
    ReceiveType    int           `gorm:"column:receive_type;type:int" json:"receiveType"`
    RegularDate    int           `gorm:"column:regular_date;type:int" json:"regularDate"`
    Transferable   int           `gorm:"column:transferable;type:int" json:"transferable"`
    UseEndTime     date.Datetime `gorm:"column:use_end_time;type:timestamp" json:"useEndTime,omitempty"`
    UseRangeType   int           `gorm:"column:use_range_type;type:int" json:"useRangeType"`
    UseStartTime   date.Datetime `gorm:"column:use_start_time;type:timestamp" json:"useStartTime,omitempty"`
}

func (CouponModels) TableName() string {
	return "coupon_models"
}
func (this CouponModels) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this CouponModels) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this CouponModels) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this CouponModels) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(CouponModels{}).Error; err != nil {
		return err
	}
	return nil
}

func NewCouponModels() CouponModels {
    return CouponModels{}
}

