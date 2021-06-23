package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.644646 +0800 CST m=+0.047963610
type CouponAccount struct {
    Contributions int            `gorm:"column:contributions;type:int" json:"contributions"`
    CreatedAt     date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt     *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    ID            uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    Mobile        string         `gorm:"column:mobile;type:varchar(255);size:1020" json:"mobile,omitempty"`
    Points        int            `gorm:"column:points;type:int" json:"points"`
    UpdatedAt     date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (CouponAccount) TableName() string {
	return "coupon_account"
}
func (this CouponAccount) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this CouponAccount) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this CouponAccount) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this CouponAccount) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(CouponAccount{}).Error; err != nil {
		return err
	}
	return nil
}

func NewCouponAccount() CouponAccount {
    return CouponAccount{}
}

