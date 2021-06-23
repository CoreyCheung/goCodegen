package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.647239 +0800 CST m=+0.050556670
type CouponDetail struct {
    CardId       string         `gorm:"column:card_id;type:varchar(255);size:1020" json:"cardId,omitempty"`
    CouponTypeId uint           `gorm:"column:coupon_type_id;type:int unsigned" json:"couponTypeId"`
    CreatedAt    date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt    *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    EndTime      date.Datetime  `gorm:"column:end_time;type:timestamp" json:"endTime,omitempty"`
    ID           uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    Mobile       string         `gorm:"not null;column:mobile;type:varchar(255);size:1020" json:"mobile,omitempty"`
    ReceiveTime  date.Datetime  `gorm:"column:receive_time;type:timestamp" json:"receiveTime,omitempty"`
    SellPoint    int            `gorm:"column:sell_point;type:int" json:"sellPoint"`
    StartTime    date.Datetime  `gorm:"column:start_time;type:timestamp" json:"startTime,omitempty"`
    Status       int            `gorm:"column:status;type:int" json:"status"`
    UpdatedAt    date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
    UseTime      date.Datetime  `gorm:"column:use_time;type:timestamp" json:"useTime,omitempty"`
}

func (CouponDetail) TableName() string {
	return "coupon_detail"
}
func (this CouponDetail) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this CouponDetail) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this CouponDetail) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this CouponDetail) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(CouponDetail{}).Error; err != nil {
		return err
	}
	return nil
}

func NewCouponDetail() CouponDetail {
    return CouponDetail{}
}

