package model

import "git.wanpinghui.com/WPH/go_common/wph"
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.741065 +0800 CST m=+0.144380114
type ThirdProductRelApplies struct {
    CertImg        string        `gorm:"column:cert_img;type:varchar(255);size:1020" json:"certImg,omitempty"`
    CreatedAt      date.Datetime `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    CustomerId     wph.Long      `gorm:"column:customer_id;type:bigint" json:"customerId"`
    ID             uint          `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    Stat           int8          `gorm:"column:stat;type:tinyint unsigned" json:"stat"`
    ThirdProductId uint          `gorm:"column:third_product_id;type:int unsigned" json:"thirdProductId"`
    UpdatedAt      date.Datetime `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (ThirdProductRelApplies) TableName() string {
	return "third_product_rel_applies"
}
func (this ThirdProductRelApplies) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this ThirdProductRelApplies) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this ThirdProductRelApplies) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this ThirdProductRelApplies) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(ThirdProductRelApplies{}).Error; err != nil {
		return err
	}
	return nil
}

func NewThirdProductRelApplies() ThirdProductRelApplies {
    return ThirdProductRelApplies{}
}

