package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.654579 +0800 CST m=+0.057896615
type CouponType struct {
    ActivitiesId      int           `gorm:"column:activities_id;type:int" json:"activitiesId"`
    BenefitType       int           `gorm:"column:benefit_type;type:int" json:"benefitType"`
    CouponName        string        `gorm:"column:coupon_name;type:varchar(255);size:1020" json:"couponName,omitempty"`
    Decrease          float64       `gorm:"column:decrease;type:double" json:"decrease"`
    Description       string        `gorm:"column:description;type:varchar(255);size:1020" json:"description,omitempty"`
    DiscountRate      float64       `gorm:"column:discount_rate;type:double" json:"discountRate"`
    EachReceiveCount  int           `gorm:"column:each_receive_count;type:int" json:"eachReceiveCount"`
    ExpiryDateType    int           `gorm:"column:expiry_date_type;type:int" json:"expiryDateType"`
    ID                uint          `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    MinCharge         float64       `gorm:"column:min_charge;type:double" json:"minCharge"`
    Points            int           `gorm:"column:points;type:int" json:"points"`
    ReceivableBalance int           `gorm:"column:receivable_balance;type:int" json:"receivableBalance"`
    ReceivableTotal   int           `gorm:"column:receivable_total;type:int" json:"receivableTotal"`
    ReceiveEndTime    date.Datetime `gorm:"column:receive_end_time;type:timestamp" json:"receiveEndTime,omitempty"`
    ReceiveStartTime  date.Datetime `gorm:"column:receive_start_time;type:timestamp" json:"receiveStartTime,omitempty"`
    ReceiveStatus     int           `gorm:"column:receive_status;type:int" json:"receiveStatus"`
    ReceiveType       int           `gorm:"column:receive_type;type:int" json:"receiveType"`
    RegularDate       int           `gorm:"column:regular_date;type:int" json:"regularDate"`
    Transferable      int           `gorm:"column:transferable;type:int" json:"transferable"`
    UseEndTime        date.Datetime `gorm:"column:use_end_time;type:timestamp" json:"useEndTime,omitempty"`
    UseRangeType      int           `gorm:"column:use_range_type;type:int" json:"useRangeType"`
    UseStartTime      date.Datetime `gorm:"column:use_start_time;type:timestamp" json:"useStartTime,omitempty"`
}

func (CouponType) TableName() string {
	return "coupon_type"
}
func (this CouponType) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this CouponType) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this CouponType) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this CouponType) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(CouponType{}).Error; err != nil {
		return err
	}
	return nil
}

func NewCouponType() CouponType {
    return CouponType{}
}

